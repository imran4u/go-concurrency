package main

import (
	"errors"
	"fmt"
)

// var ErrNotFound  = errors.New("not found")
var ErrNotFound = &ErrorNotFound{}

type ErrorNotFound struct{}

func (e *ErrorNotFound) Error() string {
	return "Not found"
}

func getUser(id int) error {
	if id == 0 {
		return fmt.Errorf("getUser failed: %w", ErrNotFound)
	}
	return nil
}

func main() {
	err := getUser(0)

	if errors.Is(err, ErrNotFound) {
		fmt.Println("User not found")
	}

	//
	fmt.Println(err.Error())
	var e *ErrorNotFound
	if errors.As(err, &e) {
		fmt.Println(e.Error())
	}

	// introduced in 1.26 
	if et, ok := errors.AsType[*ErrorNotFound](err); ok {
		fmt.Println("Et ::", et)
	}
}
