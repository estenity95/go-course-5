package task08concurrency

import "testing"

func TestIncrementConcurrently(t *testing.T) {
	const workers = 8
	const increments = 1000

	var counter Counter
	IncrementConcurrently(&counter, workers, increments)

	if got, want := counter.Value(), workers*increments; got != want {
		t.Fatalf("counter value = %d, want %d", got, want)
	}
}
