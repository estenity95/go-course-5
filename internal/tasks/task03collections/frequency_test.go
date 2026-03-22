package task03collections

import (
	"reflect"
	"testing"
)

func TestWordFrequency(t *testing.T) {
	words := []string{"go", "test", "go", "map", "test", "go"}
	want := map[string]int{
		"go":   3,
		"test": 2,
		"map":  1,
	}

	if got := WordFrequency(words); !reflect.DeepEqual(got, want) {
		t.Fatalf("WordFrequency(%v) = %v, want %v", words, got, want)
	}
}
