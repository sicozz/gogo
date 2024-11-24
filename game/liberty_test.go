package game

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test simple end calculation for liberty.
func TestEnd(t *testing.T) {
	testNameTemplate := "liberty %v"
	tests := []struct {
		origin      position
		dir         direction
		expectedEnd position
	}{
		{
			origin:      NewPosition(0, 0),
			dir:         up,
			expectedEnd: NewPosition(0, 1),
		},
		{
			origin:      NewPosition(1, 1),
			dir:         down,
			expectedEnd: NewPosition(1, 0),
		},
		{
			origin:      NewPosition(2, 2),
			dir:         left,
			expectedEnd: NewPosition(1, 2),
		},
		{
			origin:      NewPosition(3, 3),
			dir:         right,
			expectedEnd: NewPosition(4, 3),
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf(testNameTemplate, tt.dir.string()), func(t *testing.T) {
			assert.Equal(t, tt.expectedEnd, End(newLiberty(tt.origin, tt.dir)), "Failed to calculate liberty end")
		})
	}
}
