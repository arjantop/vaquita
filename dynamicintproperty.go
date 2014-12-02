package vaquita

type DynamicIntProperty struct {
	defaultValue int
	property     *DynamicProperty
}

func newDynamicIntProperty(p *DynamicProperty, defaultValue int) *DynamicIntProperty {
	return &DynamicIntProperty{
		defaultValue: defaultValue,
		property:     p,
	}
}

func (p *DynamicIntProperty) Name() string {
	return p.property.Name()
}

func (p *DynamicIntProperty) Get() int {
	return p.property.intValueWithDefault(p.defaultValue)
}
