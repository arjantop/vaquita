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

	parsedStringProperty parsedProperty
	parsedBoolProperty   parsedProperty
	parsedIntProperty    parsedProperty
}

func newDynamicProperty(name string) *DynamicProperty {
	s := newSharedString()
	return &DynamicProperty{
		name:  name,
		value: s,

		parsedStringProperty: newParsedProperty(s, "", func(v string) (interface{}, error) {
			return v, nil
		}),
		parsedBoolProperty: newParsedProperty(s, false, func(v string) (interface{}, error) {
			if v == "true" {
				return true, nil
			} else if v == "false" {
				return false, nil
			} else {
				return false, InvalidBoolValue
			}
		}),
		parsedIntProperty: newParsedProperty(s, 0, func(v string) (interface{}, error) {
			r, err := strconv.ParseInt(v, 0, 0)
			return int(r), err
		}),
	}
}

func (p *DynamicProperty) Name() string {
	return p.name
}

func (p *DynamicProperty) LastTimeChanged() time.Time {
	return p.value.lastTimeChanged()
}

func (p *DynamicProperty) stringValue() (string, error) {
	r, err := p.parsedStringProperty.value()
	return r.(string), err
}

func (p *DynamicProperty) stringValueWithDefault(d string) string {
	return p.parsedStringProperty.valueWithDefault(d).(string)
}

func (p *DynamicProperty) boolValue() (bool, error) {
	r, err := p.parsedBoolProperty.value()
	return r.(bool), err
}

func (p *DynamicProperty) boolValueWithDefault(d bool) bool {
	return p.parsedBoolProperty.valueWithDefault(d).(bool)
}

func (p *DynamicProperty) intValue() (int, error) {
	r, err := p.parsedIntProperty.value()
	return r.(int), err
}

func (p *DynamicProperty) intValueWithDefault(d int) int {
	return p.parsedIntProperty.valueWithDefault(d).(int)
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
	p.parsedStringProperty.clear()
	p.parsedBoolProperty.clear()
	p.parsedIntProperty.clear()
}

type parsedProperty struct {
	stringValue  *sharedString
	parsed       bool
	defaultValue interface{}
	parsedValue  interface{}
	err          error
	parse        func(string) (interface{}, error)
	lock         *sync.Mutex
}

func newParsedProperty(v *sharedString, def interface{}, parse func(string) (interface{}, error)) parsedProperty {
	return parsedProperty{
		stringValue:  v,
		defaultValue: def,
		parsedValue:  def,
		parse:        parse,
		lock:         new(sync.Mutex),
	}
}

func (p *parsedProperty) value() (interface{}, error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	if !p.parsed {
		if s := p.stringValue.get(); s != nil {
			p.parsedValue, p.err = p.parse(*s)
			p.parsed = true
		} else {
			return p.parsedValue, NoValue
		}
	}
	return p.parsedValue, p.err
}

func (p *parsedProperty) valueWithDefault(d interface{}) interface{} {
	v, err := p.value()
	if err != nil {
		return d
	}
	return v
}

func (p *parsedProperty) clear() {
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
