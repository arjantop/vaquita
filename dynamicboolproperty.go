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

func (p *DynamicBoolProperty) HasValue() bool {
	_, err := p.property.boolValue()
	return err == nil
}

type ChainedBoolProperty struct {
	prop BoolProperty
	next BoolProperty
}

func NewChainedBoolProperty(f *PropertyFactory, name string, next BoolProperty) BoolProperty {
	return &ChainedBoolProperty{
		prop: f.GetBoolProperty(name, false),
		next: next,
	}
}

func (p *ChainedBoolProperty) Name() string {
	if p.prop.HasValue() {
		return p.prop.Name()
	} else {
		return p.next.Name()
	}
}

func (p *ChainedBoolProperty) Get() bool {
	if p.prop.HasValue() {
		return p.prop.Get()
	} else {
		return p.next.Get()
	}
}

func (p *ChainedBoolProperty) HasValue() bool {
	if p.prop.HasValue() {
		return true
	} else {
		return p.next.HasValue()
	}
}
