package core

type Event struct {
	Type   string
	Target interface{}
}

type Evented struct {
	EventHandlers map[string][]func(Event)
}

func (e *Evented) On(eventType string, handler func(Event)) {
	if e.EventHandlers == nil {
		e.EventHandlers = make(map[string][]func(Event))
	}
	e.EventHandlers[eventType] = append(e.EventHandlers[eventType], handler)
}

func (e *Evented) Fire(eventType string, data interface{}) {
	event := Event{Type: eventType, Target: e}
	if handlers, ok := e.EventHandlers[eventType]; ok {
		for _, handler := range handlers {
			handler(event)
		}
	}
}
