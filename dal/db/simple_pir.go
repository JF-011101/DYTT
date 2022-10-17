package db

// #cgo CFLAGS: -O3 -march=native -msse4.1 -maes -mavx2 -mavx
// #include "pir.h"
import "C"
import "fmt"

type SimplePIR struct{}

func (pi *SimplePIR) Name() string {
	return "SimplePIR"
}

func (pi *SimplePIR) PickParams(N, d, n, logq uint64) Params {
	good_p := Params{}
	found := false

	// Iteratively refine p and DB dims, until find tight values
	for mod_p := uint64(2); ; mod_p += 1 {
		l, m := ApproxSquareDatabaseDims(N, d, mod_p)

		p := Params{
			n:    n,
			logq: logq,
			l:    l,
			m:    m,
		}
		p.PickParams(false, m)

		if p.p < mod_p {
			if !found {
				panic("Error; should not happen")
			}
			good_p.PrintParams()
			return good_p
		}

		good_p = p
		found = true
	}

	panic("Cannot be reached")
	return Params{}
}

func (pi *SimplePIR) GetBW(info DBinfo, p Params) {
	offline_download := float64(p.l*p.n*p.logq) / (8.0 * 1024.0)
	fmt.Printf("\t\tOffline download: %d KB\n", uint64(offline_download))

	online_upload := float64(p.m*p.logq) / (8.0 * 1024.0)
	fmt.Printf("\t\tOnline upload: %d KB\n", uint64(online_upload))

	online_download := float64(p.l*p.logq) / (8.0 * 1024.0)
	fmt.Printf("\t\tOnline download: %d KB\n", uint64(online_download))
}

func (pi *SimplePIR) Init(info DBinfo, p Params) State {
	A := MatrixRand(p.m, p.n, p.logq, 0)
	return MakeState(A)
}

func (pi *SimplePIR) Setup(DB *Database, shared State, p Params) (State, Msg) {
	A := shared.data[0]
	H := MatrixMul(DB.Data, A)

	// map the database entries to [0, p] (rather than [-p/1, p/2]) and then
	// pack the database more tightly in memory, because the online computation
	// is memory-bandwidth-bound
	DB.Data.Add(p.p / 2)
	DB.Squish()

	return MakeState(), MakeMsg(H)
}

func (pi *SimplePIR) FakeSetup(DB *Database, p Params) (State, float64) {
	offline_download := float64(p.l*p.n*uint64(p.logq)) / (8.0 * 1024.0)
	fmt.Printf("\t\tOffline download: %d KB\n", uint64(offline_download))

	// map the database entries to [0, p] (rather than [-p/1, p/2]) and then
	// pack the database more tightly in memory, because the online computation
	// is memory-bandwidth-bound
	DB.Data.Add(p.p / 2)
	DB.Squish()

	return MakeState(), offline_download
}

func (pi *SimplePIR) Query(i uint64, shared State, p Params, info DBinfo) (State, Msg) {
	A := shared.data[0]

	fmt.Print(1)
	secret := MatrixRand(p.n, 1, p.logq, 0)
	fmt.Print(2)
	err := MatrixGaussian(p.m, 1)
	fmt.Print(3)
	query := MatrixMul(A, secret)
	fmt.Print(4)
	query.MatrixAdd(err)
	fmt.Print(5)
	query.Data[i%p.m] += C.Elem(p.Delta())
	fmt.Print(6, p.m, 9, info.squishing)
	// Pad the query to match the dimensions of the compressed DB
	if p.m%info.squishing != 0 {
		query.AppendZeros(info.squishing - (p.m % info.squishing))
	}
	fmt.Print(7)

	return MakeState(secret), MakeMsg(query)
}

func (pi *SimplePIR) Answer(DB *Database, query MsgSlice, server State, shared State, p Params) Msg {
	ans := new(Matrix)
	num_queries := uint64(len(query.Data)) // number of queries in the batch of queries
	batch_sz := DB.Data.Rows / num_queries // how many rows of the database each query in the batch maps to

	last := uint64(0)
	// Run SimplePIR's answer routine for each query in the batch
	for batch, q := range query.Data {
		if batch == int(num_queries-1) {
			batch_sz = DB.Data.Rows - last
		}
		a := MatrixMulVecSub(DB.Data.RowsM(last, batch_sz),
			q.Data[0],
			p.p/2, // map the Z_p entries from [0, p] to [-p/2, p/2]
			DB.Info.basis,
			DB.Info.squishing)
		ans.Concat(a)
		last += batch_sz
	}

	return MakeMsg(ans)
}

func (pi *SimplePIR) Recover(i uint64, batch_index uint64, offline Msg, answer Msg,
	client State, p Params, info DBinfo) uint64 {
	secret := client.data[0]
	H := offline.Data[0]
	ans := answer.Data[0]

	row := i / p.m
	interm := MatrixMul(H, secret)
	ans.MatrixSub(interm)

	var vals []uint64
	// Recover each Z_p element that makes up the desired database entry
	for j := row * info.ne; j < (row+1)*info.ne; j++ {
		noised := ans.Data[j]
		denoised := p.Round(uint64(noised))
		vals = append(vals, denoised)
		//fmt.Printf("Reconstructing row %d: %d\n", j, denoised)
	}
	ans.MatrixAdd(interm)

	return ReconstructElem(vals, i, info)
}

func (pi *SimplePIR) Reset(DB *Database, p Params) {
	// Uncompress the database, and map its entries to the range [-p/2, p/2].
	DB.Unsquish()
	DB.Data.Sub(p.p / 2)
}
