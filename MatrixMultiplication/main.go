package MatrixMultiplication

import "fmt"

func multiplication(A [][]int, B [][]int) [][]int {
	n, m, k := len(A), len(A[0]), len(B[0])
	C := make([][]int, n, n)
	for idx, _ := range C {
		C[idx] = make([]int, k, k)
	}
	fmt.Println(C)
	for i := 0; i < n; i++ {
		for j := 0; j < k; j++ {
			for l := 0; l < m; l++ {
				C[i][j] += A[i][l] * B[l][j]
			}
		}
	}
	return C
}
