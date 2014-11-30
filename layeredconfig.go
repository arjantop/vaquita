package vaquita

import "sync"

type LayeredConfig struct {
	configList    []DynamicConfig
	defaultConfig DynamicConfig
	lock          *sync.RWMutex

	*EventHandler
}

func NewLayeredConfig() *LayeredConfig {
	cfgs := make([]DynamicConfig, 1)
	cfgs[0] = NewEmptyMapConfig()
	cfg := &LayeredConfig{
		configList:    cfgs,
		defaultConfig: cfgs[0],
		lock:          new(sync.RWMutex),
		EventHandler:  NewEventHandler(),
	}
	cfgs[0].Subscribe(cfg.eventHandler)
	return cfg
}

func (c *LayeredConfig) eventHandler(e *ChangeEvent) {
	handledEvent := c.handleEvent(e)
	if handledEvent != nil {
		c.Notify(handledEvent)
	}
}

func (c *LayeredConfig) handleEvent(e *ChangeEvent) *ChangeEvent {
	c.lock.RLock()
	defer c.lock.RUnlock()
	_, index, ok := c.getSourceAndIndex(e.Name())
	eventConfigIndex := c.getConfigIndex(e.Source())
	if ok && eventConfigIndex > index {
		return nil
	}
	e = NewChangeEvent(c, e.Event(), e.Name(), e.Value())
	switch e.Event() {
	case PropertyRemoved:
		if v, ok := c.getProperty(e.Name()); ok {
			e.event = PropertyChanged
			e.value = v
		}
	}
	return e
}

func (c *LayeredConfig) getSourceAndIndex(name string) (DynamicConfig, int, bool) {
	for i, cfg := range c.configList {
		if ok := cfg.HasProperty(name); ok {
			return cfg, i, true
		}
	}
	return nil, 0, false
}

func (c *LayeredConfig) getConfigIndex(target Config) int {
	for i, cfg := range c.configList {
		if CompareIdentity(target, cfg) {
			return i
		}
	}
	panic("unreachable")
}

func (c *LayeredConfig) SetProperty(name string, value string) {
	c.defaultConfig.SetProperty(name, value)
}

func (c *LayeredConfig) HasProperty(name string) bool {
	c.lock.RLock()
	defer c.lock.RUnlock()
	for _, cfg := range c.configList {
		if cfg.HasProperty(name) {
			return true
		}
	}
	return false
}

func (c *LayeredConfig) GetProperty(name string) (string, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.getProperty(name)
}

func (c *LayeredConfig) getProperty(name string) (string, bool) {
	for _, cfg := range c.configList {
		if v, ok := cfg.GetProperty(name); ok {
			return v, true
		}
	}
	return "", false
}

func (c *LayeredConfig) RemoveProperty(name string) {
	c.defaultConfig.RemoveProperty(name)
}

func (c *LayeredConfig) Add(cfg DynamicConfig) uint {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.addAtIndex(cfg, c.lastConfigIndex())
}

func (c *LayeredConfig) AddAtIndex(cfg DynamicConfig, index uint) uint {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.addAtIndex(cfg, index)
}

func (c *LayeredConfig) addAtIndex(cfg DynamicConfig, index uint) uint {
	if index > c.lastConfigIndex() {
		index = c.lastConfigIndex()
	}
	c.configList = insert(c.configList, cfg, index)
	cfg.Subscribe(c.eventHandler)
	return index
}

func (c *LayeredConfig) lastConfigIndex() uint {
	return uint(len(c.configList)) - 1
}

func insert(configList []DynamicConfig, cfg DynamicConfig, index uint) []DynamicConfig {
	configList = append(configList, nil)
	copy(configList[index+1:], configList[index:])
	configList[index] = cfg
	return configList
}
