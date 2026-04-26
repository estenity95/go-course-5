package task04quick

import (
	"slices"
	"testing"
	"testing/quick"
)

func TestReverseCopyProperties(t *testing.T) {
	property := func(values []int) bool {
		original := append([]int(nil), values...)
		reversed := ReverseCopy(values)

		return slices.Equal(values, original) && slices.Equal(ReverseCopy(reversed), original)
	}

	if err := quick.Check(property, &quick.Config{MaxCount: 50}); err != nil {
		t.Fatalf("quick.Check failed: %v", err)
	}
}

func BenchmarkReverseCopy(b *testing.B) {
	values := make([]int, 1000)
	for i := range values {
		values[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ReverseCopy(values)
	}
}
