package ddd

import (
	"context"
	"sync"
)

type EventSubscriber interface {
	Subscribe(event Event, handler EventHandler)
}

type EventPublisher interface {
	Publish(ctx context.Context, events ...Event) error
}

type EventDispatcher struct {
	handlers map[string][]EventHandler
	mu       sync.Mutex
}

var _ interface {
	EventSubscriber
	EventPublisher
} = (*EventDispatcher)(nil)

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandler),
	}
}

func (e *EventDispatcher) Publish(ctx context.Context, events ...Event) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	for _, event := range events {
		for _, handler := range e.handlers[event.EventName()] {
			if err := handler(ctx, event); err != nil {
				return err
			}
		}
	}
	return nil
}

func (e *EventDispatcher) Subscribe(event Event, handler EventHandler) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.handlers[event.EventName()] = append(e.handlers[event.EventName()], handler)
}
