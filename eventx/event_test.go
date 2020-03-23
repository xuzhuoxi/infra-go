package eventx

import (
	"fmt"
	"testing"
)

func TestEventDispatcher(t *testing.T) {
	listener := func(ed *EventData) {
		fmt.Println(ed.Data)
	}
	println(&listener, &listener)
	var dispatcher IEventDispatcher = &EventDispatcher{}
	dispatcher.AddEventListener("A", listener)
	dispatcher.DispatchEvent("A", dispatcher, 11111)
	fmt.Println("---")
	dispatcher.AddEventListener("A", listener)
	dispatcher.DispatchEvent("A", dispatcher, 222)
	fmt.Println("---")
	dispatcher.RemoveEventListener("A", listener)
	dispatcher.DispatchEvent("A", dispatcher, 333)
	fmt.Println("---")
	dispatcher.AddEventListener("A", listener)
	dispatcher.AddEventListener("A", listener)
	dispatcher.AddEventListener("A", listener)
	dispatcher.AddEventListener("A", listener)
	dispatcher.RemoveEventListenerByType("A")
	dispatcher.DispatchEvent("A", dispatcher, 333)
}
