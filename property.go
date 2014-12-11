package vaquita

import "time"

type Property interface {
	Name() string
	HasValue() bool
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

type DurationProperty interface {
	Get() time.Duration
	Property
}
