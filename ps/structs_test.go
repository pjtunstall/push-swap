package ps

import "testing"

func TestGetNumsString(t *testing.T) {
	a, _ := NewStack("1 2 3")
	res := a.GetNumsString()
	if res != "1 2 3" {
		t.Errorf("a.GetNumsString() failed: got %v, want \"1 2 3\"", res)
	}
}

func TestGetNUmsSlice(t *testing.T) {
	a, _ := NewStack("1 2 3")
	res := a.GetNumsSlice()
	if len(res) != 3 || res[0] != 1 || res[1] != 2 || res[2] != 3 {
		t.Errorf("a.GetNumsSlice() failed: got %v, want [1 2 3]", res)
	}
}

func TestCopy(t *testing.T) {
	a, _ := NewStack("1 2 3")
	b := a.Copy()
	if len(a.Nums) != len(b.Nums) {
		t.Errorf("a.Copy() failed: lenth of original: %v, copy length: %v", len(a.Nums), len(b.Nums))
	}
	if a.Top != b.Top {
		t.Errorf("a.Copy() failed: Top of original: %v, copy Top: %v", a.Top, b.Top)
	}
	for i := 0; i < len(a.Nums); i++ {
		if a.Nums[i] != b.Nums[i] {
			t.Errorf("a.Copy() failed: got %v, want %v", b, a)
		}
	}
}

func TestMax(t *testing.T) {
	a, _ := NewStack("1 2 3")
	i, M, err := a.Max()
	if i != 2 || M != 3 || err != nil {
		t.Errorf("a.Max() from [1 2 3] failed: got %v, %v, %v, want 2, 3, nil", i, M, err)
	}
}

func TestMin(t *testing.T) {
	a, _ := NewStack("1 2 3")
	i, M, err := a.Min()
	if i != 0 || M != 1 || err != nil {
		t.Errorf("a.Min() from [1 2 3] failed: got %v, %v, %v, want 0, 1, nil", i, M, err)
	}
}
