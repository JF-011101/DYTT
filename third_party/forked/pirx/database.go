package pirx

import (
	"fmt"
	"math"
)

type DBinfo struct {
	N          uint64 // number of DB entries.
	row_length uint64 // number of bits per DB entry.

	packing uint64 // number of DB entries per Z_p elem, if log(p) > DB entry size.
	ne      uint64 // number of Z_p elems per DB entry, if DB entry size > log(p).

	x uint64 // tunable param that governs communication,
	// must be in range [1, ne] and must be a divisor of ne;
	// represents the number of times the scheme is repeated.
	p    uint64 // plaintext modulus.
	logq uint64 // (logarithm of) ciphertext modulus.

	// For in-memory DB compression
	basis     uint64
	squishing uint64
	cols      uint64
}

type Database struct {
	Info DBinfo
	Data *Matrix
}

func (DB *Database) Squish() {
	fmt.Printf("Original DB dims: ")
	DB.Data.Dim()

	DB.Info.basis = 10
	DB.Info.squishing = 3
	DB.Info.cols = DB.Data.Cols
	DB.Data.Squish(DB.Info.basis, DB.Info.squishing)

	fmt.Printf("After squishing, with compression factor %d: ", DB.Info.squishing)
	DB.Data.Dim()

	// Check that params allow for this compression
	if (DB.Info.p > (1 << DB.Info.basis)) || (DB.Info.logq < DB.Info.basis*DB.Info.squishing) {
		panic("Bad params")
	}
}

func (DB *Database) Unsquish() {
	DB.Data.Unsquish(DB.Info.basis, DB.Info.squishing, DB.Info.cols)
}

// Store the database with entries decomposed into Z_p elements, and mapped to [-p/2, p/2]
// Z_p elements that encode the same database entry are stacked vertically below each other.
func ReconstructElem(vals []uint64, index uint64, info DBinfo) uint64 {
	q := uint64(1 << info.logq)

	for i, _ := range vals {
		vals[i] = (vals[i] + info.p/2) % q
		vals[i] = vals[i] % info.p
	}

	val := Reconstruct_from_base_p(info.p, vals)

	if info.packing > 0 {
		val = Base_p((1 << info.row_length), val, index%info.packing)
	}

	return val
}

func (DB *Database) GetElem(i uint64) uint64 {
	if i >= DB.Info.N {
		panic("Index out of range")
	}

	col := i % DB.Data.Cols
	row := i / DB.Data.Cols

	if DB.Info.packing > 0 {
		new_i := i / DB.Info.packing
		col = new_i % DB.Data.Cols
		row = new_i / DB.Data.Cols
	}

	var vals []uint64
	for j := row * DB.Info.ne; j < (row+1)*DB.Info.ne; j++ {
		vals = append(vals, DB.Data.Get(j, col))
	}

	return ReconstructElem(vals, i, DB.Info)
}

// Find smallest l, m such that l*m >= N*ne and ne divides l, where ne is
// the number of Z_p elements per DB entry determined by row_length and p.
func ApproxSquareDatabaseDims(N, row_length, p uint64) (uint64, uint64) {
	db_elems, elems_per_entry, _ := Num_DB_entries(N, row_length, p)
	l := uint64(math.Floor(math.Sqrt(float64(db_elems))))

	rem := l % elems_per_entry
	if rem != 0 {
		l += elems_per_entry - rem
	}

	m := uint64(math.Ceil(float64(db_elems) / float64(l)))

	return l, m
}

// Find smallest l, m such that l*m >= N*ne and ne divides l, where ne is
// the number of Z_p elements per DB entry determined by row_length and p, and m >=
// lower_bound_m.
func ApproxDatabaseDims(N, row_length, p, lower_bound_m uint64) (uint64, uint64) {
	l, m := ApproxSquareDatabaseDims(N, row_length, p)
	if m >= lower_bound_m {
		return l, m
	}

	m = lower_bound_m
	db_elems, elems_per_entry, _ := Num_DB_entries(N, row_length, p)
	l = uint64(math.Ceil(float64(db_elems) / float64(m)))

	rem := l % elems_per_entry
	if rem != 0 {
		l += elems_per_entry - rem
	}

	return l, m
}

func SetupDB(N, row_length uint64, p *Params) *Database {
	if (N == 0) || (row_length == 0) {
		panic("Empty database!")
	}

	D := new(Database)

	D.Info.N = N
	D.Info.row_length = row_length
	D.Info.p = p.p
	D.Info.logq = p.logq

	db_elems, elems_per_entry, entries_per_elem := Num_DB_entries(N, row_length, p.p)
	D.Info.ne = elems_per_entry
	D.Info.x = D.Info.ne
	D.Info.packing = entries_per_elem

	for D.Info.ne%D.Info.x != 0 {
		D.Info.x += 1
	}

	D.Info.basis = 0
	D.Info.squishing = 0

	fmt.Printf("Total packed DB size is ~%f MB\n",
		float64(p.l*p.m)*math.Log2(float64(p.p))/(1024.0*1024.0*8.0))

	if db_elems > p.l*p.m {
		panic("Params and database size don't match")
	}

	if p.l%D.Info.ne != 0 {
		panic("Number of DB elems per entry must divide DB height")
	}

	return D
}

func MakeRandomDB(N, row_length uint64, p *Params) *Database {
	D := SetupDB(N, row_length, p)
	D.Data = MatrixRand(p.l, p.m, 0, p.p)

	// Map DB elems to [-p/2; p/2]
	D.Data.Sub(p.p / 2)

	return D
}

func MakeDB(N, row_length uint64, p *Params, vals []uint64) *Database {
	D := SetupDB(N, row_length, p)
	D.Data = MatrixZeros(p.l, p.m)

	if uint64(len(vals)) != N {
		panic("Bad input DB")
	}

	if D.Info.packing > 0 {
		// Pack multiple DB elems into each Z_p elem
		at := uint64(0)
		cur := uint64(0)
		coeff := uint64(1)
		for i, elem := range vals {
			cur += (elem * coeff)
			coeff *= (1 << row_length)
			if ((i+1)%int(D.Info.packing) == 0) || (i == len(vals)-1) {
				D.Data.Set(cur, at/p.m, at%p.m)
				at += 1
				cur = 0
				coeff = 1
			}
		}
	} else {
		// Use multiple Z_p elems to represent each DB elem
		for i, elem := range vals {
			for j := uint64(0); j < D.Info.ne; j++ {
				D.Data.Set(Base_p(D.Info.p, elem, j), (uint64(i)/p.m)*D.Info.ne+j, uint64(i)%p.m)
			}
		}
	}

	// Map DB elems to [-p/2; p/2]
	D.Data.Sub(p.p / 2)

	return D
}
