package main

import (
	"errors"
)

// GreetingService provides greeting options.
type GreetingService interface {
	Greeting(string) (string, error)
}

type greetingService struct{}

func (greetingService) Greeting(s string) (string, error) {
	if s == "" {
		return "Hello, World", nil
	}
	return "Hello, " + s, nil
}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")
