package vaquita_test

import (
	"testing"

	"github.com/arjantop/vaquita"
	"github.com/stretchr/testify/assert"
)

func TestLayeredConfigSetProperty(t *testing.T) {
	c := vaquita.NewLayeredConfig()

	assert.False(t, c.HasProperty("p"))
	_, ok := c.GetProperty("p")
	assert.False(t, ok)

	c.SetProperty("p", "v")
	assertPropertyValue(t, c, "p", "v")
}

func TestLayeredConfigRemoveProperty(t *testing.T) {
	c := vaquita.NewLayeredConfig()

	c.SetProperty("p", "v")
	assert.True(t, c.HasProperty("p"))

	c.RemoveProperty("p")
	assert.False(t, c.HasProperty("p"))
}

func TestLayeredConfigAdd(t *testing.T) {
	c := vaquita.NewLayeredConfig()

	c.SetProperty("p", "v")

	c2 := vaquita.NewEmptyMapConfig()
	c2.SetProperty("p", "v2")
	assert.Equal(t, 0, c.Add(c2))

	assertPropertyValue(t, c, "p", "v2")

	c3 := vaquita.NewEmptyMapConfig()
	c3.SetProperty("p", "v3")
	assert.Equal(t, 1, c.Add(c3))

	assertPropertyValue(t, c, "p", "v2")
}

func TestLayeredConfigAddAtIndex(t *testing.T) {
	c := vaquita.NewLayeredConfig()

	c.SetProperty("p", "v")

	c2 := vaquita.NewEmptyMapConfig()
	c2.SetProperty("p", "v2")
	assert.Equal(t, 0, c.Add(c2))

	assertPropertyValue(t, c, "p", "v2")

	c3 := vaquita.NewEmptyMapConfig()
	c3.SetProperty("p", "v3")
	assert.Equal(t, 0, c.AddAtIndex(c3, 0))

	assertPropertyValue(t, c, "p", "v3")
}

func TestLayeredConfigAddAtIndexOutOfBounds(t *testing.T) {
	c := vaquita.NewLayeredConfig()

	c2 := vaquita.NewEmptyMapConfig()
	assert.Equal(t, 0, c.AddAtIndex(c2, 10))

	c3 := vaquita.NewEmptyMapConfig()
	assert.Equal(t, 1, c.AddAtIndex(c3, 10))
}

func TestLayeredConfigMultipleLayers(t *testing.T) {
	c := vaquita.NewLayeredConfig()

	c.SetProperty("p", "v")

	c2 := vaquita.NewEmptyMapConfig()
	c2.SetProperty("p", "v2")
	assert.Equal(t, 0, c.Add(c2))

	c3 := vaquita.NewEmptyMapConfig()
	c3.SetProperty("p", "v3")
	assert.Equal(t, 0, c.AddAtIndex(c3, 0))

	assertPropertyValue(t, c, "p", "v3")

	c2.RemoveProperty("p")
	assertPropertyValue(t, c, "p", "v3")

	c3.RemoveProperty("p")
	assertPropertyValue(t, c, "p", "v")

	c.RemoveProperty("p")
	assert.False(t, c.HasProperty("p"))

	c2.SetProperty("p2", "abc")
	assertPropertyValue(t, c, "p2", "abc")
}

func assertPropertyValue(t *testing.T, c vaquita.Config, name string, expected string) {
	assert.True(t, c.HasProperty(name))
	v, ok := c.GetProperty(name)
	assert.True(t, ok)
	assert.Equal(t, expected, v)
}
