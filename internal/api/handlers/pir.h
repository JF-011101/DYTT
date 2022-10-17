#include <stdint.h>
#include <stddef.h>

typedef uint32_t Elem;

void transposeCli(Elem *out, const Elem *in, size_t rows, size_t cols);

void matMulCli(Elem *out, const Elem *a, const Elem *b,
    size_t aRows, size_t aCols, size_t bCols);

void matMulVecCli(Elem *out, const Elem *a, const Elem *b,
    size_t aRows, size_t aCols);

void matMulVecSubCli(Elem *out, const Elem *a, const Elem *b,
    size_t aRows, size_t aCols, size_t sub);
