package vaquita_test

import (
	"testing"

	"github.com/arjantop/vaquita"
	"github.com/stretchr/testify/assert"
)

func TestConfigTestIdentity(t *testing.T) {
	c1 := vaquita.NewEmptyMapConfig()
	c2 := vaquita.NewEmptyMapConfig()
	c3 := vaquita.NewEmptyMapConfig()
	assert.True(t, vaquita.CompareIdentity(c1, c1))
	assert.False(t, vaquita.CompareIdentity(c1, c2))
	assert.False(t, vaquita.CompareIdentity(c1, c3))
	assert.True(t, vaquita.CompareIdentity(c2, c2))
	assert.False(t, vaquita.CompareIdentity(c2, c3))
}
