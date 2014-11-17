package vaquita

type Event int

const (
	PropertySet Event = iota
	PropertyRemoved
	PropertyChanged
)

type ChangeEvent struct {
	source      Config
	event       Event
	name, value string
}

func NewChangeEvent(source Config, event Event, name, value string) *ChangeEvent {
	return &ChangeEvent{
		source: source,
		event:  event,
		name:   name,
		value:  value,
	}
}

func (e *ChangeEvent) Source() Config {
	return e.source
}

func (e *ChangeEvent) Event() Event {
	return e.event
}

func (e *ChangeEvent) Name() string {
	return e.name
}

func (e *ChangeEvent) Value() string {
	return e.value
}

type Callback func(e *ChangeEvent)

type Observable interface {
	Subscribe(Callback)
}
