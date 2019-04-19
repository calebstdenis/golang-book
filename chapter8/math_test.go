package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		min, err := Min([]float64{-1.1, -1, 0, 1, 1.1})
		assert.Equal(t, min, -1.1)
		assert.Nil(t, err)
	})

	t.Run("empty slice", func(t *testing.T) {
		min, err := Min([]float64{})
		assert.EqualError(t, err, "Empty slice")
		assert.Empty(t, min)
	})

	t.Run("nil slice", func(t *testing.T) {
		min, err := Min(nil)
		assert.EqualError(t, err, "Empty slice")
		assert.Empty(t, min)
	})
}

// Same shit for TestMax
