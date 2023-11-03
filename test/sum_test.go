package test

import "testing"

func TestSum(t *testing.T) {
	got := sum(1, 1)

	want := 2

	if got != want {
		t.Errorf("fail to sum 1 + 1, expected %d, got %d", want, got)
	}
}
