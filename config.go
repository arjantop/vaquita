package vaquita

import "reflect"

type Config interface {
	SetProperty(name string, value string)
	HasProperty(name string) bool
	GetProperty(name string) (string, bool)
	RemoveProperty(name string)
}

type DynamicConfig interface {
	Config
	Observable
}

func CompareIdentity(cfg1 Config, cfg2 Config) bool {
	v1 := reflect.ValueOf(cfg1)
	v2 := reflect.ValueOf(cfg2)
	return v1.Pointer() == v2.Pointer()
}
