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

func (p *DynamicIntProperty) HasValue() bool {
	_, err := p.property.intValue()
	return err == nil
}

type ChainedIntProperty struct {
	prop IntProperty
	next IntProperty
}

func NewChainedIntProperty(f *PropertyFactory, name string, next IntProperty) IntProperty {
	return &ChainedIntProperty{
		prop: f.GetIntProperty(name, 0),
		next: next,
	}
}

func (p *ChainedIntProperty) Name() string {
	if p.prop.HasValue() {
		return p.prop.Name()
	} else {
		return p.next.Name()
	}
}

func (p *ChainedIntProperty) Get() int {
	if p.prop.HasValue() {
		return p.prop.Get()
	} else {
		return p.next.Get()
	}
}

func (p *ChainedIntProperty) HasValue() bool {
	if p.prop.HasValue() {
		return true
	} else {
		return p.next.HasValue()
	}
}
