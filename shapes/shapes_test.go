package shapes

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArea(t *testing.T) {
	tests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{
			name:     "Rectangle 10x5",
			shape:    Rectangle{Width: 10, Height: 5},
			expected: 50.0,
		},
		{
			name:     "Circle radius 7",
			shape:    Circle{Radius: 7},
			expected: math.Pi * 49,
		},
		{
			name:     "Rectangle 0x0",
			shape:    Rectangle{Width: 0, Height: 0},
			expected: 0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.InDelta(t, tt.expected, tt.shape.Area(), 0.01)
		})
	}
}

func TestPrintArea(t *testing.T) {
	tests := []struct {
		name     string
		shape    Shape
		expected string
	}{
		{
			name:     "Print Rectangle",
			shape:    Rectangle{Width: 4, Height: 2.5},
			expected: "Area: 10.00",
		},
		{
			name:     "Print Circle",
			shape:    Circle{Radius: 1},
			expected: "Area: 3.14",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, PrintArea(tt.shape))
		})
	}
}
