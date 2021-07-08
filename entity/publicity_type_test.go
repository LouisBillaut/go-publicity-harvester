package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPublicityType(t *testing.T) {
	pubID := NewID()
	typeID := NewID()
	pt := NewPublicityType(pubID, typeID)
	assert.Equal(t, pubID, pt.PublicityID)
	assert.Equal(t, typeID, pt.TypeID)
}
