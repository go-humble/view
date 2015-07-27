package view

import (
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

// EventListener represents a listener attached to a single type of event on
// one or more elements.
type EventListener struct {
	typ         string
	elements    []dom.Element
	listener    func(dom.Event)
	jsListeners []func(*js.Object)
}

// AddEventListener addds an event listener to one or more child elements of the
// view. selector is a query selector which starts at the view's root element.
// AddEventListener will add an event listener to *all* child elements that
// match the given selector. listener is a function that will be called when the
// event is triggered. Because of the way gopherjs works, listener cannot be a
// blocking function. See https://github.com/gopherjs/gopherjs#goroutines for
// more information. If the view is re-rendered, you may need to remove the old
// listeners and call AddEventListener again.
func AddEventListener(view View, eventType string, selector string, listener func(dom.Event)) *EventListener {
	eventListener := &EventListener{
		listener: listener,
		typ:      eventType,
	}
	eventListener.elements = view.Element().QuerySelectorAll(selector)
	for _, el := range eventListener.elements {
		jsListener := el.AddEventListener(eventType, true, listener)
		eventListener.jsListeners = append(eventListener.jsListeners, jsListener)
	}
	return eventListener
}

// Remove removes the event listener from all elements.
func (eventListener *EventListener) Remove() {
	for i, el := range eventListener.elements {
		el.RemoveEventListener(eventListener.typ, true,
			eventListener.jsListeners[i])
	}
}

// Elements returns a slice of all elements that the event listener is attached
// to.
func (eventListener *EventListener) Elements() []dom.Element {
	return eventListener.Elements()
}
