package task07time

import (
	"testing"
	"time"
)

func TestIsStoreOpen(t *testing.T) {
	tests := []struct {
		name    string
		current time.Time
		want    bool
	}{
		{
			name:    "before opening",
			current: time.Date(2026, time.March, 22, 8, 59, 0, 0, time.UTC),
			want:    false,
		},
		{
			name:    "during working hours",
			current: time.Date(2026, time.March, 22, 14, 0, 0, 0, time.UTC),
			want:    true,
		},
		{
			name:    "at closing time",
			current: time.Date(2026, time.March, 22, 18, 0, 0, 0, time.UTC),
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			now := func() time.Time {
				return tt.current
			}

			if got := IsStoreOpen(now); got != tt.want {
				t.Fatalf("IsStoreOpen() = %t, want %t", got, tt.want)
			}
		})
	}
}
