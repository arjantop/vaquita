package vaquita

type Property interface {
	Name() string
}

type StringProperty interface {
	Get() string
	Property
}

type IntProperty interface {
	Get() int
	Property
}

type BoolProperty interface {
	Get() bool
	Property
}
