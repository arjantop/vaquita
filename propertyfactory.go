package vaquita

import "sync"

type PropertyFactory struct {
	properties map[string]*DynamicProperty
	lock       *sync.Mutex
	config     Config
}

func NewPropertyFactory(owner Config) *PropertyFactory {
	return &PropertyFactory{
		properties: make(map[string]*DynamicProperty),
		lock:       new(sync.Mutex),
		config:     owner,
	}
}

func (f *PropertyFactory) GetDynamicStringProperty(name string, defaultValue string) *DynamicStringProperty {
	f.lock.Lock()
	defer f.lock.Unlock()
	p, ok := f.properties[name]
	if !ok {
		p = newDynamicProperty(name)
		if v, ok := f.config.GetProperty(name); ok {
			// If there is a value in the config that owns this property we
			// assign it before returning the property.
			p.setValue(v)
		}
		f.properties[name] = p
	}
	return newDynamicStringProperty(p, defaultValue)
}

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
