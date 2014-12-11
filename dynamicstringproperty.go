package vaquita

type DynamicStringProperty struct {
	defaultValue string
	property     *DynamicProperty
}

func newDynamicStringProperty(p *DynamicProperty, defaultValue string) *DynamicStringProperty {
	return &DynamicStringProperty{
		defaultValue: defaultValue,
		property:     p,
	}
}

func (p *DynamicStringProperty) Name() string {
	return p.property.Name()
}

func (p *DynamicStringProperty) Get() string {
	return p.property.stringValueWithDefault(p.defaultValue)
}

func (p *DynamicStringProperty) HasValue() bool {
	_, err := p.property.stringValue()
	return err == nil
}

type ChainedStringProperty struct {
	prop StringProperty
	next StringProperty
}

func NewChainedStringProperty(f *PropertyFactory, name string, next StringProperty) StringProperty {
	return &ChainedStringProperty{
		prop: f.GetStringProperty(name, ""),
		next: next,
	}
}

func (p *ChainedStringProperty) Name() string {
	if p.prop.HasValue() {
		return p.prop.Name()
	} else {
		return p.next.Name()
	}
}

func (p *ChainedStringProperty) Get() string {
	if p.prop.HasValue() {
		return p.prop.Get()
	} else {
		return p.next.Get()
	}
}

func (p *ChainedStringProperty) HasValue() bool {
	if p.prop.HasValue() {
		return true
	} else {
		return p.next.HasValue()
	}
}
