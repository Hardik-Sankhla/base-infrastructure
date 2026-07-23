package runtime

import (
	"sync"
	"time"
)

// EventType defines the type of event
type EventType string

const (
	DiscoveryStarted        EventType = "DiscoveryStarted"
	DiscoveryFinished       EventType = "DiscoveryFinished"
	DiscoveryStageStarted   EventType = "DiscoveryStageStarted"
	DiscoveryStageCompleted EventType = "DiscoveryStageCompleted"
	DiscoveryStageFailed    EventType = "DiscoveryStageFailed"
	PluginLoaded            EventType = "PluginLoaded"
	PluginInstalled         EventType = "PluginInstalled"
	PluginFailed            EventType = "PluginFailed"
	HealthChanged           EventType = "HealthChanged"
	RollbackStarted         EventType = "RollbackStarted"
	RollbackCompleted       EventType = "RollbackCompleted"
)

// Event payload
type Event struct {
	Type      EventType
	Timestamp time.Time
	Payload   interface{}
}

// StageEventPayload is the structured payload for discovery stage events
type StageEventPayload struct {
	StageName string
	Status    string
	Timestamp time.Time
	Duration  time.Duration
	Error     string
	Metadata  map[string]string
}

// Handler function signature
type Handler func(Event)

// EventBus interface
type EventBus interface {
	Subscribe(eventType EventType, handler Handler)
	Publish(eventType EventType, payload interface{})
}

type DefaultEventBus struct {
	handlers map[EventType][]Handler
	mu       sync.RWMutex
}

func NewEventBus() *DefaultEventBus {
	return &DefaultBus{
		handlers: make(map[EventType][]Handler),
	}
}

func (b *DefaultEventBus) Subscribe(eventType EventType, handler Handler) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.handlers[eventType] = append(b.handlers[eventType], handler)
}

func (b *DefaultEventBus) Publish(eventType EventType, payload interface{}) {
	b.mu.RLock()
	handlers := b.handlers[eventType]
	b.mu.RUnlock()

	event := Event{
		Type:      eventType,
		Timestamp: time.Now(),
		Payload:   payload,
	}

	// Dispatch to handlers
	for _, h := range handlers {
		go h(event) // async dispatch
	}
}
