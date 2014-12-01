package vaquita

import "sync"

type PropertyFactory struct {
	properties map[string]*DynamicProperty
	lock       *sync.RWMutex
	config     DynamicConfig
}

func NewPropertyFactory(owner DynamicConfig) *PropertyFactory {
	factory := &PropertyFactory{
		properties: make(map[string]*DynamicProperty),
		lock:       new(sync.RWMutex),
		config:     owner,
	}
	owner.Subscribe(factory.eventHandler)
	return factory
}

func (f *PropertyFactory) eventHandler(e *ChangeEvent) {
	f.lock.RLock()
	defer f.lock.RUnlock()
	if p, ok := f.properties[e.Name()]; ok {
		if e.Event() == PropertyRemoved {
			p.clear()
		} else {
			p.setValue(e.Value())
		}
	}
}

func (f *PropertyFactory) GetDynamicStringProperty(name string, defaultValue string) *DynamicStringProperty {
	f.lock.Lock()
	defer f.lock.Unlock()
	p, ok := f.properties[name]
	if !ok {
		p = newDynamicProperty(name)
		if v, ok := f.config.GetProperty(name); ok {
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
