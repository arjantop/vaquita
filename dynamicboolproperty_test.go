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

	c.SetProperty("p", "false")
	assert.Equal(t, false, p.Get())

	c.SetProperty("p", "true")
	assert.Equal(t, true, p.Get())

	c.RemoveProperty("p")
	assert.Equal(t, true, p.Get())
}

func TestDynamicBoolPropertyInvalidValue(t *testing.T) {
	c := vaquita.NewEmptyMapConfig()
	f := vaquita.NewPropertyFactory(c)

	p := f.GetBoolProperty("p", true)

	c.SetProperty("p", "fals")
	assert.Equal(t, true, p.Get(), "The value is still the default")
}
