package hll

import "testing"

func TestProbeMergeNilReturnsError(t *testing.T) {
	h, err := New(10)
	if err != nil {
		t.Fatal(err)
	}
	if err := h.Merge(nil); err == nil {
		t.Fatal("Merge(nil) returned nil error")
	}
}
