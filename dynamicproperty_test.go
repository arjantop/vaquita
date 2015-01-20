package vaquita

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPropertyName(t *testing.T) {
	p := newDynamicProperty("property.name")
	assert.Equal(t, "property.name", p.Name())
}

func TestPropertySetValue(t *testing.T) {
	p := newDynamicProperty("")

	// Default is returned if no value is set.
	assert.Equal(t, "abc", p.stringValueWithDefault("abc"))
	assert.Equal(t, true, p.boolValueWithDefault(true))
	assert.Equal(t, 123, p.intValueWithDefault(123))

	p.setValue("false")
	assert.Equal(t, "false", p.stringValueWithDefault("abc"))
	assert.Equal(t, false, p.boolValueWithDefault(true))
	assert.Equal(t, 123, p.intValueWithDefault(123))

	p.setValue("44")
	assert.Equal(t, "44", p.stringValueWithDefault("abc"))
	assert.Equal(t, true, p.boolValueWithDefault(true))
	assert.Equal(t, 44, p.intValueWithDefault(123))
}

func TestPropertySetValueSameValue(t *testing.T) {
	p := newDynamicProperty("")
	assert.True(t, p.LastTimeChanged().IsZero())
	p.setValue("abc")
	changedTime := p.LastTimeChanged()
	assert.Equal(t, "abc", p.stringValueWithDefault(""))
	p.setValue("abc")
	assert.Equal(t, "abc", p.stringValueWithDefault(""))
	assert.Equal(t, changedTime, p.LastTimeChanged())
}

func TestPropertySetValueBool(t *testing.T) {
	p := newDynamicProperty("")

	p.setValue("false")
	assert.Equal(t, false, p.boolValueWithDefault(true))
	p.setValue("true")
	assert.Equal(t, true, p.boolValueWithDefault(false))
}

func TestPropertySetValueInt(t *testing.T) {
	p := newDynamicProperty("")

	p.setValue("255")
	assert.Equal(t, 255, p.intValueWithDefault(0))
	p.setValue("0xFF")
	assert.Equal(t, 255, p.intValueWithDefault(0))
	p.setValue("0377")
	assert.Equal(t, 255, p.intValueWithDefault(0))
}

func TestPropertyClear(t *testing.T) {
	p := newDynamicProperty("")
	p.setValue("abc")
	assert.Equal(t, "abc", p.stringValueWithDefault(""))
	p.clear()
	assert.Equal(t, "foo", p.stringValueWithDefault("foo"))
}

func BenchmarkDynamicPropertyBool(b *testing.B) {
	p := newDynamicProperty("")
	for i := 0; i < b.N; i++ {
		p.boolValueWithDefault(false)
	}
}

func BenchmarkDynamicPropertyInt(b *testing.B) {
	p := newDynamicProperty("")
	for i := 0; i < b.N; i++ {
		p.intValueWithDefault(0)
	}
}
