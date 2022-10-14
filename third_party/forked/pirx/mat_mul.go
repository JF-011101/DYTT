package pirx

const (
	COMPRESSION int = 3
	BASIS       int = 10
	MASK        int = (1 << BASIS) - 1
)

type Elem = uint64
type SizeT = uint64

func matMul(out []Elem, a []Elem, b []Elem, aRows SizeT, aCols SizeT, bCols SizeT) {
	var i, k, j Elem

	for i = 0; i < aRows; i++ {
		for k = 0; k < aCols; k++ {
			for j = 0; j < bCols; j++ {
				out[bCols*i+j] += a[aCols*i+k] * b[bCols*k+j]
			}
		}
	}
}

func matMulVec(out []Elem, a []Elem, b []Elem, aRows SizeT, aCols SizeT) {
	var (
		tmp  Elem
		i, j SizeT
	)
	for i = 0; i < aRows; i++ {
		tmp = 0
		for j = 0; j < aCols; j++ {
			tmp += a[aCols*i+j] * b[j]
		}
		out[i] = tmp
	}
}

func matMulVecSub(out []Elem, a []Elem, b []Elem, aRows SizeT, aCols SizeT, sub SizeT) {
	var (
		tmp, val Elem
		i, j     SizeT
		k        int
	)
	for i = 0; i < aRows; i++ {
		tmp = 0
		for j = 0; j < aCols; j++ {
			for k = 0; k < COMPRESSION; k++ {
				val = (a[aCols*i+j] >> Elem(k*BASIS)) & Elem(MASK)
				val -= sub
				tmp += val * b[j*Elem(COMPRESSION)+Elem(k)]
			}
		}
	}

}

func transpose(out []Elem, in []Elem, rows SizeT, cols SizeT) {
	var (
		i, j SizeT
	)
	for i = 0; i < rows; i++ {
		for j = 0; j < cols; j++ {
			out[j*rows+i] = in[i*cols+j]
		}
	}
}
