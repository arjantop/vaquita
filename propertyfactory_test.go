package vaquita_test

import (
	"testing"

	"github.com/arjantop/vaquita"
	"github.com/stretchr/testify/assert"
)

func TestPropertyFactoryMultipleProperties(t *testing.T) {
	c := vaquita.NewEmptyMapConfig()
	f := vaquita.NewPropertyFactory(c)

	p := f.GetDynamicStringProperty("p", "")
	c.SetProperty("p", "v")
	p2 := f.GetDynamicStringProperty("p", "")

	assert.Equal(t, p.Get(), p2.Get())
}

func TestPropertyFactoryPropertyLocalDefault(t *testing.T) {
	c := vaquita.NewEmptyMapConfig()
	f := vaquita.NewPropertyFactory(c)

	p := f.GetDynamicStringProperty("p", "d")
	p2 := f.GetDynamicStringProperty("p", "d2")

	assert.Equal(t, "d", p.Get())
	assert.Equal(t, "d2", p2.Get())
}

func TestDynamicStringPropertyDefaultValue(t *testing.T) {
	f := vaquita.NewPropertyFactory(vaquita.NewEmptyMapConfig())
	p := f.GetDynamicStringProperty("p", "d")
	assert.Equal(t, "p", p.Name())

	assert.Equal(t, "d", p.Get())
}

func TestDynamicStringPropertyGet(t *testing.T) {
	c := vaquita.NewEmptyMapConfig()
	f := vaquita.NewPropertyFactory(c)

	c.SetProperty("p", "v")

	p := f.GetDynamicStringProperty("p", "d")
	assert.Equal(t, "v", p.Get())

	c.SetProperty("p", "v1")
	assert.Equal(t, "v1", p.Get())

	c.SetProperty("p", "v2")
	assert.Equal(t, "v2", p.Get())

	c.RemoveProperty("p")
	assert.Equal(t, "d", p.Get())
}
