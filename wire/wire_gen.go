// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

// Injectors from wire.go:

func initializeEvent(msg string) event {
	message2 := newMessage(msg)
	greeter2 := newGreeter(message2)
	event2 := newEvent(greeter2)
	return event2
}
