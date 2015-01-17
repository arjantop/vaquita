package vaquita

import (
	"errors"
	"strconv"
	"sync"
	"time"
)

var (
	NoValue          = errors.New("no property value")
	InvalidBoolValue = errors.New("invalid boolean value")
)

type DynamicProperty struct {
	name  string
	value *sharedString

	parsedBoolProperty parsedBoolProperty
	parsedIntProperty  parsedIntProperty
}

func newDynamicProperty(name string) *DynamicProperty {
	s := newSharedString()
	return &DynamicProperty{
		name:  name,
		value: s,

		parsedBoolProperty: newParsedBoolProperty(s),
		parsedIntProperty:  newParsedIntProperty(s),
	}
}

func (p *DynamicProperty) Name() string {
	return p.name
}

func (p *DynamicProperty) LastTimeChanged() time.Time {
	return p.value.lastTimeChanged()
}

func (p *DynamicProperty) stringValue() (string, error) {
	r := p.value.get()
	if r == nil {
		return "", NoValue
	}
	return *r, nil
}

func (p *DynamicProperty) stringValueWithDefault(d string) string {
	r := p.value.get()
	if r == nil {
		return d
	}
	return *r
}

func (p *DynamicProperty) boolValue() (bool, error) {
	return p.parsedBoolProperty.value()
}

func (p *DynamicProperty) boolValueWithDefault(d bool) bool {
	return p.parsedBoolProperty.valueWithDefault(d)
}

func (p *DynamicProperty) intValue() (int, error) {
	return p.parsedIntProperty.value()
}

func (p *DynamicProperty) intValueWithDefault(d int) int {
	return p.parsedIntProperty.valueWithDefault(d)
}

func (p *DynamicProperty) setValue(v string) {
	p.value.withValue(func(s **string) {
		if *s != nil && **s == v {
			// If the value did not change do nothing.
			return
		}
		*s = &v
		p.clearParsedProperties()
	})
}

func (p *DynamicProperty) clear() {
	p.value.withValue(func(s **string) {
		*s = nil
		p.clearParsedProperties()
	})
}

func (p *DynamicProperty) clearParsedProperties() {
	p.parsedBoolProperty.clear()
	p.parsedIntProperty.clear()
}

type parsedBoolProperty struct {
	stringValue *sharedString
	parsed      bool
	parsedValue bool
	err         error
	lock        sync.Mutex
}

func newParsedBoolProperty(v *sharedString) parsedBoolProperty {
	return parsedBoolProperty{
		stringValue: v,
	}
}

func (p *parsedBoolProperty) value() (bool, error) {
	p.lock.Lock()
	if !p.parsed {
		if s := p.stringValue.get(); s != nil {
			if *s == "true" {
				p.parsedValue = true
			} else if *s == "false" {
				p.parsedValue = false
			} else {
				p.err = InvalidBoolValue
			}
		} else {
			p.err = NoValue
		}
		p.parsed = true
	}
	val, err := p.parsedValue, p.err
	p.lock.Unlock()
	return val, err
}

func (p *parsedBoolProperty) valueWithDefault(d bool) bool {
	v, err := p.value()
	if err != nil {
		return d
	}
	return v
}

func (p *parsedBoolProperty) clear() {
	p.lock.Lock()
	p.parsed = false
	p.err = nil
	p.lock.Unlock()
}

type parsedIntProperty struct {
	stringValue *sharedString
	parsed      bool
	parsedValue int
	err         error
	lock        sync.Mutex
}

func newParsedIntProperty(v *sharedString) parsedIntProperty {
	return parsedIntProperty{
		stringValue: v,
	}
}

func (p *parsedIntProperty) value() (int, error) {
	p.lock.Lock()
	if !p.parsed {
		if s := p.stringValue.get(); s != nil {
			r, err := strconv.ParseInt(*s, 0, 0)
			p.parsedValue = int(r)
			p.err = err
		} else {
			p.err = NoValue
		}
		p.parsed = true
	}
	val, err := p.parsedValue, p.err
	p.lock.Unlock()
	return val, err
}

func (p *parsedIntProperty) valueWithDefault(d int) int {
	v, err := p.value()
	if err != nil {
		return d
	}
	return v
}

func (p *parsedIntProperty) clear() {
	p.lock.Lock()
	p.parsed = false
	p.err = nil
	p.lock.Unlock()
}

type sharedString struct {
	value       *string
	timeChanged time.Time
	lock        sync.RWMutex
}

func newSharedString() *sharedString {
	return &sharedString{}
}

func (s *sharedString) withValue(f func(**string)) {
	s.lock.Lock()
	prev := s.value
	f(&s.value)
	if prev != s.value {
		s.timeChanged = time.Now()
	}
	s.lock.Unlock()
}

func (s *sharedString) lastTimeChanged() time.Time {
	s.lock.RLock()
	t := s.timeChanged
	s.lock.RUnlock()
	return t
}

func (s *sharedString) get() *string {
	s.lock.RLock()
	v := s.value
	s.lock.RUnlock()
	return v
}
