package vaquita_test

import (
	"testing"

	"github.com/arjantop/vaquita"
	"github.com/stretchr/testify/assert"
)

func TestGetPropertyEmpty(t *testing.T) {
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

func AssertDynamicConfigEvents(t *testing.T, c vaquita.DynamicConfig) {
	var eventCount int

	// Removing a nonexistent property.
	s := c.Subscribe(func(e *vaquita.ChangeEvent) {
		eventCount++
	})
	c.RemoveProperty("p")
	assert.Equal(t, 0, eventCount)
	c.Unsubscribe(s)

	// Setting a property.
	s = c.Subscribe(func(e *vaquita.ChangeEvent) {
		eventCount++
		assert.Equal(t, "p2", e.Name())
		assert.Equal(t, "v", e.Value())
		assert.Equal(t, vaquita.PropertySet, e.Event())
		assert.Equal(t, c, e.Source())
	})
	c.SetProperty("p2", "v")
	assert.Equal(t, 1, eventCount)
	c.Unsubscribe(s)

	// Changing a propery value.
	s = c.Subscribe(func(e *vaquita.ChangeEvent) {
		eventCount++
		assert.Equal(t, "p2", e.Name())
		assert.Equal(t, "v2", e.Value())
		assert.Equal(t, vaquita.PropertyChanged, e.Event())
		assert.Equal(t, c, e.Source())
	})
	c.SetProperty("p2", "v2")
	assert.Equal(t, 2, eventCount)
	c.Unsubscribe(s)

	// Changing a propery value.
	s = c.Subscribe(func(e *vaquita.ChangeEvent) {
		eventCount++
		assert.Equal(t, "p2", e.Name())
		assert.Equal(t, "", e.Value())
		assert.Equal(t, vaquita.PropertyRemoved, e.Event())
		assert.Equal(t, c, e.Source())
	})
	c.RemoveProperty("p2")
	assert.Equal(t, 3, eventCount)
	c.Unsubscribe(s)
}

func TestDynamicConfigEvents(t *testing.T) {
	c := vaquita.NewEmptyMapConfig()
	AssertDynamicConfigEvents(t, c)
}
