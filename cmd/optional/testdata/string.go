// Code generated by go generate. DO NOT EDIT.
// This file was generated by robots at 2020-04-10 22:05:29.582843 +0000 UTC

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
	*s = MakeString(v)
}

// Unset unsets the string value.
func (s *String) Unset() {
	*s = String{}
}

// Present returns whether or not the value is present.
func (s String) Present() bool {
	return s.set
}

// Get returns the string value or panics if not present.
func (s String) Get() string {
	if !s.Present() {
		panic("value not present")
	}
	return s.val
}

// GetOr returns the string value or a default value if not present.
func (s String) GetOr(v string) string {
	if s.Present() {
		return s.val
	}
	return v
}

// GetOrBool returns the string value or false if not present.
func (s String) GetOrBool() (string, bool) {
	if !s.Present() {
		var zero string
		return zero, false
	}
	return s.val, true
}

// GetOrErr returns the string value or an error if not present.
func (s String) GetOrErr() (string, error) {
	if !s.Present() {
		var zero string
		return zero, errors.New("value not present")
	}
	return s.val, nil
}

// If calls the function fn with the value if the value is present.
func (s String) If(fn func(string)) {
	if s.Present() {
		fn(s.val)
	}
}

// Map applies the function fn to the contained value (if any) and returns a new
// option value.
func (s String) Map(fn func(string) string) String {
	if s.Present() {
		return MakeString(fn(s.val))
	}
	return s
}

// And returns an empty String option value if not present, otherwise returns
// optb.
func (s String) And(optb String) String {
	if s.Present() {
		return optb
	}
	return String{}
}

// Or returns the String option value if present, otherwise returns optb.
func (s String) Or(optb String) String {
	if s.Present() {
		return s
	}
	return optb
}

// MarshalJSON implements the json.Marshaler interface.
func (s String) MarshalJSON() ([]byte, error) {
	if s.Present() {
		return json.Marshal(s.val)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
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
