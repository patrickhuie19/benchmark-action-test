package main

import (
	"testing"
	"time"
)

// BenchmarkAdd is a dummy benchmark test to monitor for regressions in overal benchmark CI behavior.
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(1, 2)
		time.Sleep(1 * time.Nanosecond)
	}

}
