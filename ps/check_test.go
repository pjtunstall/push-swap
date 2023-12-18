package ps

import (
	"testing"
)

func TestCheck(t *testing.T) {
	s := "0 1 2 3 4 5 6 7 8 9"
	a, _ := NewStack(s)
	b, _ := NewStack("")
	rotatable, sorted := Check(a, b)
	if rotatable != true {
		t.Errorf("Expected %v to be rotatable", s)
	}
	if sorted != true {
		t.Errorf("Expected %v to be sorted", s)
	}

	s = "3 4 5 6 7 8 9 0 1 2"
	a, _ = NewStack(s)
	b, _ = NewStack("")
	rotatable, sorted = Check(a, b)
	if rotatable != true {
		t.Errorf("Expected %v to be rotatable", s)
	}
	if sorted != false {
		t.Errorf("Expected %v not to be sorted", s)
	}

	s = "1 0 2 3 4 5 6 7 8 9"
	a, _ = NewStack(s)
	b, _ = NewStack("")
	rotatable, sorted = Check(a, b)
	if rotatable == true {
		t.Errorf("Did not expect %v to be rotatable", s)
	}
	if sorted != false {
		t.Errorf("Did not expect %v to be sorted", s)
	}
}
