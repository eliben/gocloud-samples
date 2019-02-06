//+build wireinject

package main

import "github.com/google/wire"

func InitializeEvent(s string) (Event, error) {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}, nil
}
