package ps

import "testing"

func TestRun(t *testing.T) {
	a, _ := NewStack("1 2 3")
	b, _ := NewStack("4 5 6")
	instructions := []string{"ra", "rra", "rb", "rrb", "rr", "rrr", "pa", "pb", "sa", "sb", "ss"}
	Run(&a, &b, instructions)
	A := a.GetNumsSlice()
	B := b.GetNumsSlice()
	failLength := len(A) != 3 || len(B) != 3
	failA := A[0] != 1 || A[1] != 2 || A[2] != 3
	failB := B[0] != 4 || B[1] != 5 || B[2] != 6
	if failLength || failA || failB {
		t.Errorf("Run(A=1 2 3, B=4 5 6, instructions) failed: got A = %v, B = %v; want A = [1 2 3], B = [4 5 6]", A, B)
	}
}
