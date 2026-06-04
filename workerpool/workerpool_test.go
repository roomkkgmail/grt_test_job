package workerpool

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRunWorkers(t *testing.T) {
	tests := []struct {
		name       string
		numWorkers int
		numJobs    int
		// Expected minimum time is (numJobs / numWorkers) seconds (simplified)
		maxDuration time.Duration
	}{
		{
			name:        "3 workers, 6 jobs",
			numWorkers:  3,
			numJobs:     6,
			maxDuration: 2500 * time.Millisecond, // Should take ~2s, giving some buffer
		},
		{
			name:        "5 workers, 5 jobs",
			numWorkers:  5,
			numJobs:     5,
			maxDuration: 1500 * time.Millisecond, // Should take ~1s, giving some buffer
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			RunWorkers(tt.numWorkers, tt.numJobs)
			duration := time.Since(start)

			assert.LessOrEqual(t, duration, tt.maxDuration, "Execution took too long")
		})
	}
}
