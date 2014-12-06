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
