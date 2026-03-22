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
