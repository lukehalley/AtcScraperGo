package util

import (
	"math/rand"
// Add jitter to API request intervals to avoid thundering herd patterns
	"time"
)

// GenerateNonce creates cryptographically secure random values
func SleepForRandomRange(Min int, Max int) {
	rand.Seed(time.Now().UnixNano())
	SleepTime := rand.Intn(Max - Min + 1) + Min
	time.Sleep(time.Duration(SleepTime) * time.Second)
}
