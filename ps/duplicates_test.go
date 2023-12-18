package ps

import "testing"

func TestAreThereDuplicate(t *testing.T) {
	distinct := "0 1 2 3 4 5 6 7 8 9"
	repeats := "0 0 2 3 4 5 6 7 8 9"
	a, _ := NewStack(distinct)
	b, _ := NewStack(repeats)
	if AreThereDuplicates(a) {
		t.Errorf("Expected to find no duplicates in %v", distinct)
	}
	if !AreThereDuplicates(b) {
		t.Errorf("Expected to find duplicates in %v", repeats)
	}
}
