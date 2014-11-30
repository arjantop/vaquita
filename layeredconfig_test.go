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

func TestLayeredConfigDynamicConfigEvents(t *testing.T) {
	c := vaquita.NewLayeredConfig()
	AssertDynamicConfigEvents(t, c)
}

func TestLayeredConfigDynamicConfigEventsMultipleLayers(t *testing.T) {
	c := vaquita.NewLayeredConfig()

	c.SetProperty("p", "v")

	c2 := vaquita.NewEmptyMapConfig()
	c.Add(c2)

	c2.SetProperty("p", "v2")

	var eventCount int

	// Removing property in non-primary config.
	s := c.Subscribe(func(e *vaquita.ChangeEvent) {
		eventCount++
	})
	c.RemoveProperty("p")
	assert.Equal(t, 0, eventCount)
	c.Unsubscribe(s)

	// Setting value of non-primary config.
	s = c.Subscribe(func(e *vaquita.ChangeEvent) {
		eventCount++
	})
	c.SetProperty("p", "v3")
	assert.Equal(t, 0, eventCount)
	c.Unsubscribe(s)

	// Setting a value of primary config.
	s = c.Subscribe(func(e *vaquita.ChangeEvent) {
		eventCount++
		assert.Equal(t, "p", e.Name())
		assert.Equal(t, "v4", e.Value())
		assert.Equal(t, vaquita.PropertyChanged, e.Event())
		assert.Equal(t, c, e.Source())
	})
	c2.SetProperty("p", "v4")
	assert.Equal(t, 1, eventCount)
	c.Unsubscribe(s)

	// Removing a property from primary config.
	s = c.Subscribe(func(e *vaquita.ChangeEvent) {
		eventCount++
		assert.Equal(t, "p", e.Name())
		assert.Equal(t, "v3", e.Value())
		assert.Equal(t, vaquita.PropertyChanged, e.Event())
		assert.Equal(t, c, e.Source())
	})
	c2.RemoveProperty("p")
	assert.Equal(t, 2, eventCount)
	c.Unsubscribe(s)

	// Removing a property from primary config.
	s = c.Subscribe(func(e *vaquita.ChangeEvent) {
		eventCount++
		assert.Equal(t, "p", e.Name())
		assert.Equal(t, "", e.Value())
		assert.Equal(t, vaquita.PropertyRemoved, e.Event())
		assert.Equal(t, c, e.Source())
	})
	c.RemoveProperty("p")
	assert.Equal(t, 3, eventCount)
	c.Unsubscribe(s)
}
