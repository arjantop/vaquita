package vaquita_test

import (
	"testing"

	"github.com/arjantop/vaquita"
	"github.com/stretchr/testify/assert"
)

func TestDynamicIntPropertyDefaultValue(t *testing.T) {
	f := vaquita.NewPropertyFactory(vaquita.NewEmptyMapConfig())
	p := f.GetDynamicIntProperty("p", 1337)
	assert.Equal(t, "p", p.Name())

	assert.Equal(t, 1337, p.Get())
}

func TestDynamicIntPropertyGet(t *testing.T) {
	c := vaquita.NewEmptyMapConfig()
	f := vaquita.NewPropertyFactory(c)

	c.SetProperty("p", "123")

	p := f.GetDynamicIntProperty("p", 1)
	assert.Equal(t, 123, p.Get())

	c.SetProperty("p", "-10")
	assert.Equal(t, -10, p.Get())

	c.SetProperty("p", "9000")
	assert.Equal(t, 9000, p.Get())

	c.RemoveProperty("p")
	assert.Equal(t, 1, p.Get())
}

func TestDynamicIntPropertyInvalidValue(t *testing.T) {
	c := vaquita.NewEmptyMapConfig()
	f := vaquita.NewPropertyFactory(c)

	p := f.GetDynamicIntProperty("p", 0)

	c.SetProperty("p", "6.6")
	assert.Equal(t, 0, p.Get())

	c.SetProperty("p", "abc")
	assert.Equal(t, 0, p.Get())

	c.SetProperty("p", "10")
	assert.Equal(t, 10, p.Get())
}
