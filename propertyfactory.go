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

func (f *PropertyFactory) GetDynamicStringProperty(name, defaultValue string) *DynamicStringProperty {
	p := f.getProperty(name)
	return newDynamicStringProperty(p, defaultValue)
}

func (f *PropertyFactory) GetDynamicBoolProperty(name string, defaultValue bool) *DynamicBoolProperty {
	p := f.getProperty(name)
	return newDynamicBoolProperty(p, defaultValue)
}

func (f *PropertyFactory) GetDynamicIntProperty(name string, defaultValue int) *DynamicIntProperty {
	p := f.getProperty(name)
	return newDynamicIntProperty(p, defaultValue)
}

func (f *PropertyFactory) getProperty(name string) *DynamicProperty {
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
	return p
}
