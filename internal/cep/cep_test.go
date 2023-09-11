package cep_test

import (
	"errors"
	"github.com/igoramorim/go-practice-cep-race/internal/cep"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("Build a CEP", func(t *testing.T) {
		result, err := cep.New("04313110")
		assert.NoError(t, err)
		assert.Equal(t, "04313", result.Prefix())
		assert.Equal(t, "110", result.Suffix())
	})

	t.Run("Return error: invalid number of digits", func(t *testing.T) {
		result, err := cep.New("0431311")
		assert.True(t, errors.Is(err, cep.ErrInvalidLen))
		assert.Empty(t, result)
	})

	t.Run("Return error: contains letters", func(t *testing.T) {
		result, err := cep.New("0A313110")
		assert.True(t, errors.Is(err, cep.ErrInvalidChar))
		assert.Empty(t, result)
	})
}
