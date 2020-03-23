package eventx

import (
	"fmt"
	"testing"
)

type subDispatcher struct {
	EventDispatcher
}

func newSubDispatcher() *subDispatcher {
	return &subDispatcher{EventDispatcher{}}
}

func printListener(ed *EventData) {
	fmt.Println(ed.Data)
}

func TestEventDispatcher(t *testing.T) {
	var dispatcher = NewEventDispatcher()
	dispatcher.AddEventListener("A", printListener)
	dispatcher.DispatchEvent("A", dispatcher, 11111)
	fmt.Println("---")
	dispatcher.AddEventListener("A", printListener)
	dispatcher.DispatchEvent("A", dispatcher, 222)
	fmt.Println("---")
	dispatcher.RemoveEventListener("A", printListener)
	dispatcher.DispatchEvent("A", dispatcher, 333)
	fmt.Println("---")
	dispatcher.AddEventListener("A", printListener)
	dispatcher.AddEventListener("A", printListener)
	dispatcher.AddEventListener("A", printListener)
	dispatcher.AddEventListener("A", printListener)
	dispatcher.RemoveEventListenerByType("A")
	dispatcher.DispatchEvent("A", dispatcher, 333)
}

func TestEventDispatcherSub(t *testing.T) {
	var dispatcher = newSubDispatcher()
	dispatcher.AddEventListener("A", printListener)
	dispatcher.DispatchEvent("A", dispatcher, 11111)
	fmt.Println("---")
	dispatcher.AddEventListener("A", printListener)
	dispatcher.DispatchEvent("A", dispatcher, 222)
	fmt.Println("---")
	dispatcher.RemoveEventListener("A", printListener)
	dispatcher.DispatchEvent("A", dispatcher, 333)
	fmt.Println("---")
	dispatcher.AddEventListener("A", printListener)
	dispatcher.AddEventListener("A", printListener)
	dispatcher.AddEventListener("A", printListener)
	dispatcher.AddEventListener("A", printListener)
	dispatcher.RemoveEventListenerByType("A")
	dispatcher.DispatchEvent("A", dispatcher, 333)
}
