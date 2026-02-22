package event

const (
	EventMemoVisited = "memo.visited"
)

type Event struct {
	Type string
	Data any
}

func NewEventBus() *EventBus {
	return &EventBus{bus: make(chan Event)}
}

type EventBus struct {
	bus chan Event
}

func (e *EventBus) Publish(event Event) {
	e.bus <- event
}

func (e *EventBus) Subscribe() <-chan Event {
	return e.bus
}
