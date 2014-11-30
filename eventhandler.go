package vaquita

import "sync"

type EventHandler struct {
	subscribers    map[subscription]Callback
	lock           *sync.RWMutex
	subscriptionId int
}

func NewEventHandler() *EventHandler {
	return &EventHandler{
		subscribers:    make(map[subscription]Callback),
		lock:           new(sync.RWMutex),
		subscriptionId: 0,
	}
}

func (h *EventHandler) Notify(e *ChangeEvent) {
	h.lock.RLock()
	defer h.lock.RUnlock()
	for _, f := range h.subscribers {
		f(e)
	}
}

func (h *EventHandler) Subscribe(f Callback) subscription {
	h.lock.Lock()
	defer h.lock.Unlock()
	s := newSubscription(h.subscriptionId, h)

	h.subscribers[s] = f
	h.subscriptionId++
	return s
}

func (h *EventHandler) Unsubscribe(s subscription) error {
	h.lock.Lock()
	defer h.lock.Unlock()
	if _, ok := h.subscribers[s]; !ok {
		return InvalidSubscribtion
	}
	delete(h.subscribers, s)
	return nil
}
