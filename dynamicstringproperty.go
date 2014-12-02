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
