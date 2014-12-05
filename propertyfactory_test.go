package vaquita_test

import (
	"testing"

	"github.com/arjantop/vaquita"
	"github.com/stretchr/testify/assert"
)

func TestPropertyFactoryMultipleProperties(t *testing.T) {
	c := vaquita.NewEmptyMapConfig()
	f := vaquita.NewPropertyFactory(c)

	p := f.GetStringProperty("p", "")
	c.SetProperty("p", "v")
	p2 := f.GetStringProperty("p", "")

	assert.Equal(t, p.Get(), p2.Get())
}

func TestPropertyFactoryPropertyLocalDefault(t *testing.T) {
	c := vaquita.NewEmptyMapConfig()
	f := vaquita.NewPropertyFactory(c)

	p := f.GetStringProperty("p", "d")
	p2 := f.GetStringProperty("p", "d2")

	assert.Equal(t, "d", p.Get())
	assert.Equal(t, "d2", p2.Get())
}
