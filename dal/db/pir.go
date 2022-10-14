package db

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"runtime/pprof"
	"time"
)

// Defines the interface for PIR with preprocessing schemes
type PIR interface {
	Name() string

	PickParams(N, d, n, logq uint64) Params

	GetBW(info DBinfo, p Params)

	Init(info DBinfo, p Params) State

	Setup(DB *Database, shared State, p Params) (State, Msg)
	FakeSetup(DB *Database, p Params) (State, float64) // used for benchmarking online phase

	Query(i uint64, shared State, p Params, info DBinfo) (State, Msg)

	Answer(DB *Database, query MsgSlice, server State, shared State, p Params) Msg

	Recover(i uint64, batch_index uint64, offline Msg, answer Msg, client State,
		p Params, info DBinfo) uint64

	Reset(DB *Database, p Params) // reset DB to its correct state, if modified during execution
}

// Run PIR's online phase, with a random preprocessing (to skip the offline phase).
// Gives accurate bandwidth and online time measurements.
func RunFakePIR(pi PIR, DB *Database, p Params, i []uint64,
	f *os.File, profile bool) (float64, float64, float64, float64) {
	fmt.Printf("Executing %s\n", pi.Name())
	debug.SetGCPercent(-1)

	num_queries := uint64(len(i))
	if DB.Data.Rows/num_queries < DB.Info.ne {
		panic("Too many queries to handle!")
	}
	shared_state := pi.Init(DB.Info, p)

	fmt.Println("Setup...")
	server_state, bw := pi.FakeSetup(DB, p)
	offline_comm := bw
	runtime.GC()

	fmt.Println("Building query...")
	start := time.Now()
	var query MsgSlice
	for index, _ := range i {
		_, q := pi.Query(i[index], shared_state, p, DB.Info)
		query.data = append(query.data, q)
	}
	printTime(start)
	online_comm := float64(query.size() * uint64(p.logq) / (8.0 * 1024.0))
	fmt.Printf("\t\tOnline upload: %f KB\n", online_comm)
	bw += online_comm
	runtime.GC()

	fmt.Println("Answering query...")
	if profile {
		pprof.StartCPUProfile(f)
	}
	start = time.Now()
	answer := pi.Answer(DB, query, server_state, shared_state, p)
	elapsed := printTime(start)
	if profile {
		pprof.StopCPUProfile()
	}
	rate := printRate(p, elapsed, len(i))
	online_down := float64(answer.size() * uint64(p.logq) / (8.0 * 1024.0))
	fmt.Printf("\t\tOnline download: %f KB\n", online_down)
	bw += online_down
	online_comm += online_down

	runtime.GC()
	debug.SetGCPercent(100)
	pi.Reset(DB, p)

	if offline_comm+online_comm != bw {
		panic("Should not happen!")
	}

	return rate, bw, offline_comm, online_comm
}

// Run full PIR scheme (offline + online phases).
func RunPIR(pi PIR, DB *Database, p Params, i []uint64) (float64, float64) {
	fmt.Printf("Executing %s\n", pi.Name())
	debug.SetGCPercent(-1)

	num_queries := uint64(len(i))
	if DB.Data.Rows/num_queries < DB.Info.ne {
		panic("Too many queries to handle!")
	}
	batch_sz := DB.Data.Rows / (DB.Info.ne * num_queries) * DB.Data.Cols
	bw := float64(0)

	shared_state := pi.Init(DB.Info, p)

	fmt.Println("Setup...")
	start := time.Now()
	server_state, offline_download := pi.Setup(DB, shared_state, p)
	printTime(start)
	comm := float64(offline_download.size() * uint64(p.logq) / (8.0 * 1024.0))
	fmt.Printf("\t\tOffline download: %f KB\n", comm)
	bw += comm
	runtime.GC()

	fmt.Println("Building query...")
	start = time.Now()
	var client_state []State
	var query MsgSlice
	for index, _ := range i {
		index_to_query := i[index] + uint64(index)*batch_sz
		cs, q := pi.Query(index_to_query, shared_state, p, DB.Info)
		client_state = append(client_state, cs)
		query.data = append(query.data, q)
	}
	runtime.GC()
	printTime(start)
	comm = float64(query.size() * uint64(p.logq) / (8.0 * 1024.0))
	fmt.Printf("\t\tOnline upload: %f KB\n", comm)
	bw += comm
	runtime.GC()

	fmt.Println("Answering query...")
	start = time.Now()
	answer := pi.Answer(DB, query, server_state, shared_state, p)
	elapsed := printTime(start)
	rate := printRate(p, elapsed, len(i))
	comm = float64(answer.size() * uint64(p.logq) / (8.0 * 1024.0))
	fmt.Printf("\t\tOnline download: %f KB\n", comm)
	bw += comm
	runtime.GC()

	pi.Reset(DB, p)
	fmt.Println("Reconstructing...")
	start = time.Now()

	for index, _ := range i {
		index_to_query := i[index] + uint64(index)*batch_sz
		val := pi.Recover(index_to_query, uint64(index), offline_download, answer,
			client_state[index], p, DB.Info)

		if DB.GetElem(index_to_query) != val {
			fmt.Printf("Batch %d (querying index %d -- row should be >= %d): Got %d instead of %d\n",
				index, index_to_query, DB.Data.Rows/4, val, DB.GetElem(index_to_query))
			panic("Reconstruct failed!")
		}
	}
	fmt.Println("Success!")
	printTime(start)

	runtime.GC()
	debug.SetGCPercent(100)
	return rate, bw
}
