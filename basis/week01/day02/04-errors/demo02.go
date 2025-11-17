package main

import (
	"errors"
	"fmt"
)

type wrapError struct {
	msg string
	err error
}

func (e *wrapError) Error() string {
	return e.msg
}

func (e *wrapError) Unwrap() error {
	return e.err
}

func main() {
	err := errors.New("the is a raw error")
	wrapErr := fmt.Errorf("error %w", err)
	fmt.Println(wrapErr)
}
