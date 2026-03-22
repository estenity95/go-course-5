package task07time

import "time"

func IsStoreOpen(now func() time.Time) bool {
	hour := now().Hour()
	return hour >= 9 && hour < 18
}
