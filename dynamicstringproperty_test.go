package vaquita_test

import (
	"testing"

	"github.com/arjantop/vaquita"
	"github.com/stretchr/testify/assert"
)

func TestDynamicStringPropertyDefaultValue(t *testing.T) {
	f := vaquita.NewPropertyFactory(vaquita.NewEmptyMapConfig())
	p := f.GetStringProperty("p", "d")
	assert.Equal(t, "p", p.Name())

	assert.Equal(t, "d", p.Get())
}

func TestDynamicStringPropertyGet(t *testing.T) {
	c := vaquita.NewEmptyMapConfig()
	f := vaquita.NewPropertyFactory(c)

	c.SetProperty("p", "v")

	p := f.GetStringProperty("p", "d")
	assert.Equal(t, "v", p.Get())
	assert.True(t, p.HasValue())

	c.SetProperty("p", "v1")
	assert.Equal(t, "v1", p.Get())
	assert.True(t, p.HasValue())

	c.SetProperty("p", "v2")
	assert.Equal(t, "v2", p.Get())
	assert.True(t, p.HasValue())

	c.RemoveProperty("p")
	assert.Equal(t, "d", p.Get())
	assert.False(t, p.HasValue())
}

func TestChainedStringProperty(t *testing.T) {
	c := vaquita.NewEmptyMapConfig()
	f := vaquita.NewPropertyFactory(c)

	p := vaquita.NewChainedStringProperty(f, "a", f.GetStringProperty("b", "db"))

	assert.Equal(t, "db", p.Get())
	assert.Equal(t, "b", p.Name())
	assert.False(t, p.HasValue())

	c.SetProperty("b", "vb")
	assert.Equal(t, "vb", p.Get())
	assert.Equal(t, "b", p.Name())
	assert.True(t, p.HasValue())

	c.SetProperty("a", "va")
	assert.Equal(t, "va", p.Get())
	assert.Equal(t, "a", p.Name())
	assert.True(t, p.HasValue())
}
