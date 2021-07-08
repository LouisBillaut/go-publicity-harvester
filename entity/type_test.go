package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewType(t *testing.T) {
	newType, err := NewType("test")
	assert.Nil(t, err)
	assert.Equal(t, newType.Type, "test")
	assert.NotNil(t, newType.ID)
	assert.NotEqual(t, newType.ID.String(), "00000000-0000-0000-0000-000000000000")
}
