package vaquita_test

import (
	"testing"

	"github.com/arjantop/vaquita"
	"github.com/stretchr/testify/assert"
)

func TestDynamicBoolPropertyDefaultValue(t *testing.T) {
	f := vaquita.NewPropertyFactory(vaquita.NewEmptyMapConfig())
	p := f.GetBoolProperty("p", false)
	assert.Equal(t, "p", p.Name())

	assert.Equal(t, false, p.Get())
}

func TestDynamicBoolPropertyGet(t *testing.T) {
	c := vaquita.NewEmptyMapConfig()
	f := vaquita.NewPropertyFactory(c)

	p := f.GetBoolProperty("p", true)
	assert.Equal(t, true, p.Get())
	assert.False(t, p.HasValue())

	c.SetProperty("p", "false")
	assert.Equal(t, false, p.Get())
	assert.True(t, p.HasValue())

	c.SetProperty("p", "true")
	assert.Equal(t, true, p.Get())
	assert.True(t, p.HasValue())

	c.RemoveProperty("p")
	assert.Equal(t, true, p.Get())
	assert.False(t, p.HasValue())
}

func TestDynamicBoolPropertyInvalidValue(t *testing.T) {
	c := vaquita.NewEmptyMapConfig()
	f := vaquita.NewPropertyFactory(c)

	p := f.GetBoolProperty("p", true)

	c.SetProperty("p", "fals")
	assert.Equal(t, true, p.Get(), "The value is still the default")
}

func TestChainedBoolProperty(t *testing.T) {
	c := vaquita.NewEmptyMapConfig()
	f := vaquita.NewPropertyFactory(c)

	p := vaquita.NewChainedBoolProperty(f, "a", f.GetBoolProperty("b", true))

	assert.Equal(t, true, p.Get())
	assert.Equal(t, "b", p.Name())
	assert.False(t, p.HasValue())

	c.SetProperty("b", "false")
	assert.Equal(t, false, p.Get())
	assert.Equal(t, "b", p.Name())
	assert.True(t, p.HasValue())

	c.SetProperty("a", "true")
	assert.Equal(t, true, p.Get())
	assert.Equal(t, "a", p.Name())
	assert.True(t, p.HasValue())
	c.SetProperty("a", "false")
	assert.Equal(t, false, p.Get())
	assert.True(t, p.HasValue())
}
