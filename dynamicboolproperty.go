package vaquita

type DynamicBoolProperty struct {
	defaultValue bool
	property     *DynamicProperty
}

func newDynamicBoolProperty(p *DynamicProperty, defaultValue bool) *DynamicBoolProperty {
	return &DynamicBoolProperty{
		defaultValue: defaultValue,
		property:     p,
	}
}

func (p *DynamicBoolProperty) Name() string {
	return p.property.Name()
}

func (p *DynamicBoolProperty) Get() bool {
	return p.property.boolValueWithDefault(p.defaultValue)
}
