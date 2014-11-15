package vaquita

import "sync"

type LayeredConfig struct {
	configList    []Config
	defaultConfig Config
	lock          *sync.RWMutex
}

func NewLayeredConfig() *LayeredConfig {
	cfgs := make([]Config, 1)
	cfgs[0] = NewEmptyMapConfig()
	return &LayeredConfig{
		configList:    cfgs,
		defaultConfig: cfgs[0],
		lock:          new(sync.RWMutex),
	}
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

func (c *LayeredConfig) Add(cfg Config) uint {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.addAtIndex(cfg, c.lastConfigIndex())
}

func (c *LayeredConfig) AddAtIndex(cfg Config, index uint) uint {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.addAtIndex(cfg, index)
}

func (c *LayeredConfig) addAtIndex(cfg Config, index uint) uint {
	if index > c.lastConfigIndex() {
		index = c.lastConfigIndex()
	}
	c.configList = insert(c.configList, cfg, index)
	return index
}

func (c *LayeredConfig) lastConfigIndex() uint {
	return uint(len(c.configList)) - 1
}

func insert(configList []Config, cfg Config, index uint) []Config {
	configList = append(configList, nil)
	copy(configList[index+1:], configList[index:])
	configList[index] = cfg
	return configList
}
