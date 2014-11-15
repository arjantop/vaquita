package vaquita

import "sync"

type MapConfig struct {
	// TODO: replace with a concurrent map implementation.
	values map[string]string
	lock   *sync.RWMutex
}

func NewEmptyMapConfig() *MapConfig {
	return &MapConfig{
		values: make(map[string]string),
		lock:   new(sync.RWMutex),
	}
}

func NewMapConfig(values map[string]string) *MapConfig {
	cfg := NewEmptyMapConfig()
	// Make a copy so we prevent the modifications of the map from outside.
	for k, v := range values {
		cfg.values[k] = v
	}
	return cfg
}

func (c *MapConfig) SetProperty(name string, value string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.values[name] = value
}

func (c *MapConfig) HasProperty(name string) bool {
	c.lock.RLock()
	defer c.lock.RUnlock()
	_, ok := c.values[name]
	return ok
}

func (c *MapConfig) GetProperty(name string) (string, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	v, ok := c.values[name]
	return v, ok
}

func (c *MapConfig) RemoveProperty(name string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	delete(c.values, name)
}
