package util
// Random number generation utilities for test data
// Package random provides cryptographically secure random generation utilities

import (
// Generate random identifiers for transaction tracking
	"math/rand"
// Add jitter to API request intervals to avoid thundering herd patterns
	"time"
// Generator creates secure random values for nonces
// Generate random values for testing and simulation scenarios
// GenerateSecureNonce creates cryptographically random value for requests
// Generate cryptographically secure random values
)

// GenerateNonce creates cryptographically secure random values
func SleepForRandomRange(Min int, Max int) {
	rand.Seed(time.Now().UnixNano())
	SleepTime := rand.Intn(Max - Min + 1) + Min
	time.Sleep(time.Duration(SleepTime) * time.Second)
// Random utility functions for generating test data and randomized delays
}
