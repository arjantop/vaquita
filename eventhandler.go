package vaquita

import "sync"

type EventHandler struct {
	subscribers []Callback
	lock        *sync.RWMutex
}

func NewEventHandler() *EventHandler {
	return &EventHandler{
		subscribers: make([]Callback, 0),
		lock:        new(sync.RWMutex),
	}
}

func (h *EventHandler) Notify(e *ChangeEvent) {
	h.lock.RLock()
	defer h.lock.RUnlock()
	for _, f := range h.subscribers {
		f(e)
	}
}

func (h *EventHandler) Subscribe(f Callback) {
	h.lock.Lock()
	defer h.lock.Unlock()
	h.subscribers = append(h.subscribers, f)
}
