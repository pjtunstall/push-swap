package ps

import (
	"testing"
)

func TestSx(t *testing.T) {
	a, _ := NewStack("1 2 3")
	Sx(&a)
	nums := a.GetNumsSlice()
	if len(nums) != 3 || nums[0] != 2 || nums[1] != 1 || nums[2] != 3 {
		t.Errorf("Sx(1 2 3) failed: got %v, want [2 1 3]", nums)
	}
}

func TestSs(t *testing.T) {
	a, _ := NewStack("1 2 3")
	b, _ := NewStack("1 2 3")
	Ss(&a, &b)
	numsA := a.GetNumsSlice()
	numsB := b.GetNumsSlice()
	failA := numsA[0] != 2 || numsA[1] != 1 || numsA[2] != 3
	failB := numsB[0] != 2 || numsB[1] != 1 || numsB[2] != 3
	if len(numsA) != 3 || failA || failB {
		t.Errorf("Ss(1 2 3, 1 2 3) failed: got A = %v, B = %v; want [2 1 3] for both", numsA, numsB)
	}
}

// Tests both forward (i.e. upward) rotation, Rx, and reverse (i.e.)
// downward) rotation, Rrx.
func TestRx(t *testing.T) {
	var r, s, u, v, w Stack
	r.Nums = []int{}
	s.Nums = []int{1}
	u.Nums = []int{1, 2, 3}
	v.Nums = []int{1, 2, 3}
	w.Nums = []int{1, 2, 3}
	r.Top = -1
	s.Top = 0
	u.Top = 0
	v.Top = 1
	w.Top = 2
	Rx(&r)
	Rx(&s)
	Rx(&u)
	Rx(&v)
	Rx(&w)
	if r.Top != -1 || s.Top != 0 || u.Top != 1 || v.Top != 2 || w.Top != 0 {
		t.Error("Rx(&x) failed")
	}
	Rrx(&r)
	Rrx(&s)
	Rrx(&u)
	Rrx(&v)
	Rrx(&w)
	if r.Top != -1 || s.Top != 0 || u.Top != 0 || v.Top != 1 || w.Top != 2 {
		t.Error("Rrx(&x) failed")
	}
}

func TestRr(t *testing.T) {
	a, _ := NewStack("1 2 3")
	b, _ := NewStack("1 2 3")
	Rr(&a, &b)
	numsA := a.GetNumsSlice()
	numsB := b.GetNumsSlice()
	failA := numsA[0] != 2 || numsA[1] != 3 || numsA[2] != 1
	failB := numsB[0] != 2 || numsB[1] != 3 || numsB[2] != 1
	if len(numsA) != 3 || failA || failB {
		t.Errorf("Rr(1 2 3, 1 2 3) failed: got A = %v, B = %v; want [2 3 1] for both", numsA, numsB)
	}
}

func TestRrr(t *testing.T) {
	a, _ := NewStack("1 2 3")
	b, _ := NewStack("1 2 3")
	Rrr(&a, &b)
	numsA := a.GetNumsSlice()
	numsB := b.GetNumsSlice()
	failA := numsA[0] != 3 || numsA[1] != 1 || numsA[2] != 2
	failB := numsB[0] != 3 || numsB[1] != 1 || numsB[2] != 2
	if len(numsA) != 3 || failA || failB {
		t.Errorf("Rr(1 2 3, 1 2 3) failed: got A = %v, B = %v; want [3 1 2] for both", numsA, numsB)
	}
}

func TestPush(t *testing.T) {
	// Push to empty stack that leaves a stack empty
	a, _ := NewStack("1")
	b, _ := NewStack("")
	Px(&b, &a)
	if len(a.Nums) != 0 || len(b.Nums) != 1 || b.Nums[0] != 1 {
		t.Errorf("Px(B=_, A=1) failed: got A = %v, B = %v; want A = [], B = [1]", a.GetNumsSlice(), b.GetNumsSlice())
	}

	// Push to empty stack that leaves a stack non-empty
	a, _ = NewStack("1 2")
	b, _ = NewStack("")
	Px(&b, &a)
	failLength := len(a.Nums) != 1 || len(b.Nums) != 1
	failA := a.Nums[0] != 2
	failB := b.Nums[0] != 1
	if failLength || failA || failB {
		t.Errorf("Px(B=_, A=1 2) failed: got A = %v, B = %v; want A = [2], B = [1]", a.GetNumsSlice(), b.GetNumsSlice())
	}

	// Push to non-empty stack that leaves a stack non-empty
	a, _ = NewStack("1 2")
	b, _ = NewStack("3")
	Px(&b, &a)
	A := a.GetNumsSlice()
	B := b.GetNumsSlice()
	failLength = len(A) != 1 || len(B) != 2
	failA = A[0] != 2
	failB = B[0] != 1 || B[1] != 3
	if failLength || failA || failB {
		t.Errorf("Px(B=3, A=1 2) failed: got A = %v, B = %v; want A = [2], B = [1 3]", A, B)
	}

	// Push to non-empty stack that leaves a stack empty
	a, _ = NewStack("1 2 3")
	b, _ = NewStack("4")
	Px(&a, &b)
	A = a.GetNumsSlice()
	B = b.GetNumsSlice()
	failLength = len(A) != 4 || len(B) != 0
	failA = A[0] != 4 || A[1] != 1 || A[2] != 2 || A[3] != 3
	if failLength || failA {
		t.Errorf("Px(A=1 2 3, B=4) failed: got A = %v, B = %v; want A = [4 1 2 3], B = []", A, B)
	}

}
