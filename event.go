package xcgui

// import (
//     "fmt"
// )

import (
    "github.com/codyguo/xcgui/xc"
)

type EventHandler func()

type Event struct {
    hEle     xc.HELE
    handlers []EventHandler
}

func (e *Event) Attach(handler EventHandler) int {

    for i, h := range e.handlers {
        if h == nil {
            e.handlers[i] = handler
            return i
        }

    }

    e.handlers = append(e.handlers, handler)
    return len(e.handlers) - 1
}

func (e *Event) Detach(handle int) {
    e.handlers[handle] = nil
}

type EventPublisher struct {
    event Event
}

func (p *EventPublisher) Event(hEle xc.HELE) *Event {
    p.event.hEle = hEle
    return &p.event
}

func (p *EventPublisher) Publish() {
    for _, handler := range p.event.handlers {
        if handler != nil {
            handler()
        }
    }
}
