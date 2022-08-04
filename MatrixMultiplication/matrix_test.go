package MatrixMultiplication

import (
	"reflect"
	"testing"
)

func TestMultiplication(t *testing.T) {
	A := [][]int{
		{1, 2, 3}, {3, 1, 2},
	}
	B := [][]int{
		{1, 2}, {3, 2}, {1, 2},
	}
	C := [][]int{
		{10, 12}, {8, 12},
	}
	if !reflect.DeepEqual(multiplication(A, B), C) {
		t.Errorf("Multiplication() = %v got, %v want",multiplication(A, B), C)
	}
}
