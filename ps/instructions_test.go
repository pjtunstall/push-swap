package ps

import (
	"testing"
)

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
	Rx(&u)
	Rx(&v)
	Rx(&w)
	if r.Top != -1 || s.Top != 0 || u.Top != 1 || v.Top != 2 || w.Top != 0 {
		t.Error("Rx(&x) failed")
	}
}
