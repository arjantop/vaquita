package vaquita

type Config interface {
	SetProperty(name string, value string)
	HasProperty(name string) bool
	GetProperty(name string) (string, bool)
	RemoveProperty(name string)
}
