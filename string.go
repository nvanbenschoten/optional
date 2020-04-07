// Code generated by go generate
// This file was generated by robots at 2020-04-07 21:28:58.420324 +0000 UTC

package optional

import (
	"encoding/json"
	"errors"
)

// String is an optional string.
type String struct {
	val string
	set bool
}

// MakeString creates an optional.String from a string.
func MakeString(v string) String {
	return String{val: v, set: true}
}

// Set sets the string value.
func (s *String) Set(v string) {
	*s = String{val: v, set: true}
}

// Unset unsets the string value.
func (s *String) Unset() {
	*s = String{}
}

// Get returns the string value or an error if not present.
func (s String) Get() (string, error) {
	if !s.Present() {
		var zero string
		return zero, errors.New("value not present")
	}
	return s.val, nil
}

// Present returns whether or not the value is present.
func (s String) Present() bool {
	return s.set
}

// OrElse returns the string value or a default value if the value is not present.
func (s String) OrElse(v string) string {
	if s.Present() {
		return s.val
	}
	return v
}

// If calls the function f with the value if the value is present.
func (s String) If(fn func(string)) {
	if s.Present() {
		fn(s.val)
	}
}

func (s String) MarshalJSON() ([]byte, error) {
	if s.Present() {
		return json.Marshal(s.val)
	}
	return json.Marshal(nil)
}

func (s *String) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		s.Unset()
		return nil
	}

	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	s.Set(value)
	return nil
}
