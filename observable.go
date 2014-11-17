package vaquita

import "errors"

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

type subscription struct {
	id    int
	owner Observable
}

func newSubscription(id int, owner Observable) subscription {
	return subscription{
		id:    id,
		owner: owner,
	}
}

var InvalidSubscribtion = errors.New("invalid subscription")

type Observable interface {
	Subscribe(Callback) subscription
	Unsubscribe(subscription) error
}
