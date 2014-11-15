package vaquita_test

import (
	"testing"

	"github.com/arjantop/vaquita"
	"github.com/stretchr/testify/assert"
)

func TestSetPropertyEmpty(t *testing.T) {
	c := vaquita.NewEmptyMapConfig()
	assert.False(t, c.HasProperty("p"))
	_, ok := c.GetProperty("p")
	assert.False(t, ok)
}

func TestSetProperty(t *testing.T) {
	c := vaquita.NewEmptyMapConfig()
	c.SetProperty("p", "v")
	assert.True(t, c.HasProperty("p"))
	v, ok := c.GetProperty("p")
	assert.True(t, ok)
	assert.Equal(t, "v", v)
}

func TestSetPropertyMultipleTimes(t *testing.T) {
	c := vaquita.NewEmptyMapConfig()
	c.SetProperty("p", "v")
	c.SetProperty("p", "v2")
	v, ok := c.GetProperty("p")
	assert.True(t, ok)
	assert.Equal(t, "v2", v)
}

func TestRemovePropertyProperty(t *testing.T) {
	c := vaquita.NewEmptyMapConfig()
	c.SetProperty("p", "v")
	c.RemoveProperty("p")
	assert.False(t, c.HasProperty("p"))
}

func TestRemovePropertyNonexistent(t *testing.T) {
	c := vaquita.NewEmptyMapConfig()
	c.RemoveProperty("p")
	assert.False(t, c.HasProperty("p"))
}

func TestNewMapConfig(t *testing.T) {
	m := make(map[string]string)
	m["p"] = "v"
	c := vaquita.NewMapConfig(m)
	assert.True(t, c.HasProperty("p"))

	// External modification should not be possible.
	delete(m, "p")
	m["p2"] = "v2"
	assert.True(t, c.HasProperty("p"))
	assert.False(t, c.HasProperty("p2"))
}
