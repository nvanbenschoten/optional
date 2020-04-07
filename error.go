// Code generated by go generate. DO NOT EDIT.
// This file was generated by robots at 2020-04-07 21:49:11.585155 +0000 UTC

package optional

import (
	"encoding/json"
	"errors"
)

// Error is an optional error.
type Error struct {
	val error
	set bool
}

// MakeError creates an optional.Error from a error.
func MakeError(v error) Error {
	return Error{val: v, set: true}
}

// Set sets the error value.
func (e *Error) Set(v error) {
	*e = Error{val: v, set: true}
}

// Unset unsets the error value.
func (e *Error) Unset() {
	*e = Error{}
}

// Present returns whether or not the value is present.
func (e Error) Present() bool {
	return e.set
}

// Get returns the error value or an error if not present.
func (e Error) Get() (error, error) {
	if !e.Present() {
		var zero error
		return zero, errors.New("value not present")
	}
	return e.val, nil
}

// MustGet returns the error value or panics if not present.
func (e Error) MustGet() error {
	if !e.Present() {
		panic("value not present")
	}
	return e.val
}

// OrElse returns the error value or a default value if the value is not present.
func (e Error) OrElse(v error) error {
	if e.Present() {
		return e.val
	}
	return v
}

// If calls the function f with the value if the value is present.
func (e Error) If(fn func(error)) {
	if e.Present() {
		fn(e.val)
	}
}

// MarshalJSON implements the json.Marshaler interface.
func (e Error) MarshalJSON() ([]byte, error) {
	if e.Present() {
		return json.Marshal(e.val)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (e *Error) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		e.Unset()
		return nil
	}

	var value error
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	e.Set(value)
	return nil
}
