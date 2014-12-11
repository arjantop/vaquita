package vaquita

import "time"

type DynamicDurationProperty struct {
	defaultValue time.Duration
	unit         time.Duration
	property     *DynamicProperty
}

func newDynamicDurationProperty(p *DynamicProperty, defaultValue time.Duration, unit time.Duration) *DynamicDurationProperty {
	return &DynamicDurationProperty{
		defaultValue: defaultValue,
		property:     p,
		unit:         unit,
	}
}

func (p *DynamicDurationProperty) Name() string {
	return p.property.Name()
}

func (p *DynamicDurationProperty) Get() time.Duration {
	return time.Duration(p.property.intValueWithDefault(int(p.defaultValue/p.unit))) * p.unit
}

func (p *DynamicDurationProperty) HasValue() bool {
	_, err := p.property.intValue()
	return err == nil
}

type ChainedDurationProperty struct {
	prop DurationProperty
	next DurationProperty
}

func NewChainedDurationProperty(f *PropertyFactory, name string, unit time.Duration, next DurationProperty) DurationProperty {
	return &ChainedDurationProperty{
		prop: f.GetDurationProperty(name, 0, unit),
		next: next,
	}
}

func (p *ChainedDurationProperty) Name() string {
	if p.prop.HasValue() {
		return p.prop.Name()
	} else {
		return p.next.Name()
	}
}

func (p *ChainedDurationProperty) Get() time.Duration {
	if p.prop.HasValue() {
		return p.prop.Get()
	} else {
		return p.next.Get()
	}
}

func (p *ChainedDurationProperty) HasValue() bool {
	if p.prop.HasValue() {
		return true
	} else {
		return p.next.HasValue()
	}
}
