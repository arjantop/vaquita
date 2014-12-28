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

		parsedBoolProperty: newParsedBoolProperty(s, false),
		parsedIntProperty:  newParsedIntProperty(s, 0),
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
	stringValue  *sharedString
	parsed       bool
	defaultValue bool
	parsedValue  bool
	err          error
	lock         *sync.Mutex
}

func newParsedBoolProperty(v *sharedString, def bool) parsedBoolProperty {
	return parsedBoolProperty{
		stringValue:  v,
		defaultValue: def,
		parsedValue:  def,
		lock:         new(sync.Mutex),
	}
}

func (p *parsedBoolProperty) value() (bool, error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	if !p.parsed {
		if s := p.stringValue.get(); s != nil {
			if *s == "true" {
				p.parsedValue = true
			} else if *s == "false" {
				p.parsedValue = false
			} else {
				p.err = InvalidBoolValue
			}
			p.parsed = true
		} else {
			return p.parsedValue, NoValue
		}
	}
	return p.parsedValue, p.err
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
	defer p.lock.Unlock()
	p.parsed = false
	p.parsedValue = p.defaultValue
	p.err = nil
}

type parsedIntProperty struct {
	stringValue  *sharedString
	parsed       bool
	defaultValue int
	parsedValue  int
	err          error
	lock         *sync.Mutex
}

func newParsedIntProperty(v *sharedString, def int) parsedIntProperty {
	return parsedIntProperty{
		stringValue:  v,
		defaultValue: def,
		parsedValue:  def,
		lock:         new(sync.Mutex),
	}
}

func (p *parsedIntProperty) value() (int, error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	if !p.parsed {
		if s := p.stringValue.get(); s != nil {
			r, err := strconv.ParseInt(*s, 0, 0)
			p.parsedValue = int(r)
			p.err = err
			p.parsed = true
		} else {
			return p.parsedValue, NoValue
		}
	}
	return p.parsedValue, p.err
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
	defer p.lock.Unlock()
	p.parsed = false
	p.parsedValue = p.defaultValue
	p.err = nil
}

type sharedString struct {
	value       *string
	timeChanged time.Time
	lock        *sync.RWMutex
}

func newSharedString() *sharedString {
	return &sharedString{
		lock: new(sync.RWMutex),
	}
}

func (s *sharedString) withValue(f func(**string)) {
	s.lock.Lock()
	defer s.lock.Unlock()
	prev := s.value
	f(&s.value)
	if prev != s.value {
		s.timeChanged = time.Now()
	}
}

func (s *sharedString) lastTimeChanged() time.Time {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.timeChanged
}

func (s *sharedString) get() *string {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.value
}
