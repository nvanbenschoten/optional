// Code generated by go generate. DO NOT EDIT.
// This file was generated by robots at 2020-04-10 23:14:18.512608 +0000 UTC

package optional

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
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
	*e = MakeError(v)
}

// Unset unsets the error value.
func (e *Error) Unset() {
	*e = Error{}
}

// Present returns whether or not the value is present.
func (e Error) Present() bool {
	return e.set
}

// Get returns the error value or panics if not present.
func (e Error) Get() error {
	if !e.Present() {
		panic("value not present")
	}
	return e.val
}

// GetOr returns the error value or a default value if not present.
func (e Error) GetOr(v error) error {
	if e.Present() {
		return e.val
	}
	return v
}

// GetOrBool returns the error value or false if not present.
func (e Error) GetOrBool() (error, bool) {
	if !e.Present() {
		var zero error
		return zero, false
	}
	return e.val, true
}

// GetOrErr returns the error value or an error if not present.
func (e Error) GetOrErr() (error, error) {
	if !e.Present() {
		var zero error
		return zero, errors.New("value not present")
	}
	return e.val, nil
}

// If calls the function fn with the value if the value is present.
func (e Error) If(fn func(error)) {
	if e.Present() {
		fn(e.val)
	}
}

// Map applies the function fn to the contained value (if any) and returns a new
// option value.
func (e Error) Map(fn func(error) error) Error {
	if e.Present() {
		return MakeError(fn(e.val))
	}
	return e
}

// And returns an empty Error option value if not present, otherwise returns
// optb.
func (e Error) And(optb Error) Error {
	if e.Present() {
		return optb
	}
	return Error{}
}

// Or returns the Error option value if present, otherwise returns optb.
func (e Error) Or(optb Error) Error {
	if e.Present() {
		return e
	}
	return optb
}

// Format implements the fmt.Formatter interface.
func (e Error) Format(fmtS fmt.State, verb rune) {
	if !e.Present() {
		io.WriteString(fmtS, "none")
		return
	}
	var format string
	switch verb {
	case 'v':
		if fmtS.Flag('+') {
			format = "some(%+v)"
		} else {
			format = "some(%v)"
		}
	case 's':
		format = "some(%v)"
	case 'q':
		format = "\"some(%v)\""
	}
	fmt.Fprintf(fmtS, format, e.val)
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
