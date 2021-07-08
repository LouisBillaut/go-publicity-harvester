package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPublicity(t *testing.T) {
	p := NewPublicity("visible", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36", "localhost")
	assert.Equal(t, p.Ua, "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36")
	assert.Equal(t, p.Ip, "localhost")
	assert.Equal(t, p.Type, "visible")
	assert.NotNil(t, p.ID)
	assert.NotEqual(t, p.ID.String(), "00000000-0000-0000-0000-000000000000")
	assert.NotNil(t, p.CreatedAt)
}
