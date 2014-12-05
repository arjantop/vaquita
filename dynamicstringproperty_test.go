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

	c.SetProperty("p", "v1")
	assert.Equal(t, "v1", p.Get())

	c.SetProperty("p", "v2")
	assert.Equal(t, "v2", p.Get())

	c.RemoveProperty("p")
	assert.Equal(t, "d", p.Get())
}
