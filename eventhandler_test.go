package vaquita_test

import (
	"testing"

	"github.com/arjantop/vaquita"
	"github.com/stretchr/testify/assert"
)

func TestEventHandlerNotifyNoSubscribers(t *testing.T) {
	h := vaquita.NewEventHandler()
	assert.NotPanics(t, func() {
		h.Notify(vaquita.NewChangeEvent(nil, vaquita.PropertySet, "", ""))
	})
}

func TestEventHandlerSubscribe(t *testing.T) {
	c := vaquita.NewEmptyMapConfig()
	h := vaquita.NewEventHandler()
	var event *vaquita.ChangeEvent
	h.Subscribe(func(e *vaquita.ChangeEvent) {
		event = e
	})
	h.Notify(vaquita.NewChangeEvent(c, vaquita.PropertySet, "p", "v"))
	assertEventEquals(t, event, "p", "v", vaquita.PropertySet, c)

	c2 := vaquita.NewEmptyMapConfig()
	h.Notify(vaquita.NewChangeEvent(c2, vaquita.PropertyChanged, "p2", "v2"))
	assertEventEquals(t, event, "p2", "v2", vaquita.PropertyChanged, c2)
}

func assertEventEquals(
	t *testing.T,
	event *vaquita.ChangeEvent,
	name, value string,
	ev vaquita.Event,
	c vaquita.Config) {

	assert.Equal(t, name, event.Name())
	assert.Equal(t, value, event.Value())
	assert.Equal(t, ev, event.Event())
	assert.True(t, vaquita.CompareIdentity(c, event.Source()))
}

func TestEventHandlerUnsubscribe(t *testing.T) {
	h := vaquita.NewEventHandler()
	var event *vaquita.ChangeEvent
	f := func(e *vaquita.ChangeEvent) {
		event = e
	}
	s := h.Subscribe(f)
	assert.NoError(t, h.Unsubscribe(s))
	h.Notify(vaquita.NewChangeEvent(nil, vaquita.PropertySet, "", ""))
	assert.Nil(t, event)
}

func TestEventHandlerUnsubscribeInOtherHandler(t *testing.T) {
	h := vaquita.NewEventHandler()
	h2 := vaquita.NewEventHandler()
	s := h.Subscribe(func(e *vaquita.ChangeEvent) {})
	s2 := h2.Subscribe(func(e *vaquita.ChangeEvent) {})
	assert.NotEqual(t, s, s2)
	assert.Equal(t, vaquita.InvalidSubscribtion, h2.Unsubscribe(s))
}
