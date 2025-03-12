package shared

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewId_有効なId(t *testing.T) {
	id, err := NewId(1)
	assert.NoError(t, err)
	assert.Equal(t, id.Value(), 1)
}

func TestNewId_無効なId(t *testing.T) {
	_, err := NewId(0)
	assert.Error(t, err)
}
