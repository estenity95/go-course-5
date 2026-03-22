package task01basic

import "testing"

func TestAdd(t *testing.T) {
	if got, expected := Add(2, 3), 5; got != expected {
		t.Fatalf("Add(2, 3) = %d, want %d", got, expected)
	}
}
