package vaquita_test

import (
	"testing"
	"time"

	"github.com/arjantop/vaquita"
	"github.com/stretchr/testify/assert"
)

func TestDynamicDurationPropertyDefaultValue(t *testing.T) {
	f := vaquita.NewPropertyFactory(vaquita.NewEmptyMapConfig())
	p := f.GetDurationProperty("p", 10*time.Second, time.Millisecond)
	assert.Equal(t, "p", p.Name())

	assert.Equal(t, 10*time.Second, p.Get())
}

func TestDynamicDurationPropertyGet(t *testing.T) {
	c := vaquita.NewEmptyMapConfig()
	f := vaquita.NewPropertyFactory(c)

	c.SetProperty("p", "100")

	p := f.GetDurationProperty("p", time.Second, time.Millisecond)
	assert.Equal(t, 100*time.Millisecond, p.Get())

	c.SetProperty("p", "-10")
	assert.Equal(t, -10*time.Millisecond, p.Get())

	c.SetProperty("p", "9000")
	assert.Equal(t, 9*time.Second, p.Get())

	c.RemoveProperty("p")
	assert.Equal(t, 1000*time.Millisecond, p.Get())
}

func TestDynamicDurationPropertyInvalidValue(t *testing.T) {
	c := vaquita.NewEmptyMapConfig()
	f := vaquita.NewPropertyFactory(c)

	p := f.GetDurationProperty("p", 2*time.Second, time.Second)

	c.SetProperty("p", "6.6")
	assert.Equal(t, 2*time.Second, p.Get())

	c.SetProperty("p", "abc")
	assert.Equal(t, 2*time.Second, p.Get())

	c.SetProperty("p", "10")
	assert.Equal(t, 10*time.Second, p.Get())
}
