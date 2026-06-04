package safecounter

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSafeCounter(t *testing.T) {
	tests := []struct {
		name          string
		numGoroutines int
		incPerRoutine int
	}{
		{
			name:          "1,000 routines, 1 inc each",
			numGoroutines: 1000,
			incPerRoutine: 1,
		},
		{
			name:          "100 routines, 100 inc each",
			numGoroutines: 100,
			incPerRoutine: 100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := SafeCounter{}
			var wg sync.WaitGroup

			for i := 0; i < tt.numGoroutines; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					for j := 0; j < tt.incPerRoutine; j++ {
						c.Inc()
					}
				}()
			}

			wg.Wait()

			expected := tt.numGoroutines * tt.incPerRoutine
			assert.Equal(t, expected, c.Value(), "Counter value mismatch")
		})
	}
}
