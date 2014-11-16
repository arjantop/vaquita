package vaquita_test

import (
	"testing"

	"github.com/arjantop/vaquita"
	"github.com/stretchr/testify/assert"
)

func TestDynamicStringPropertyName(t *testing.T) {
	f := vaquita.NewPropertyFactory(vaquita.NewEmptyMapConfig())
	p := f.GetDynamicStringProperty("prop.name", "")
	assert.Equal(t, "prop.name", p.Name())
}

func TestDynamicStringPropertyGetDefault(t *testing.T) {
	f := vaquita.NewPropertyFactory(vaquita.NewEmptyMapConfig())
	p := f.GetDynamicStringProperty("p", "v")
	assert.Equal(t, "v", p.Get())
}

func TestDynamicStringPropertyGetConfigValue(t *testing.T) {
	c := vaquita.NewEmptyMapConfig()
	c.SetProperty("p", "setvalue")
	f := vaquita.NewPropertyFactory(c)
	p := f.GetDynamicStringProperty("p", "v")
	assert.Equal(t, "setvalue", p.Get())
}
