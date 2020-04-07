// Code generated by go generate. DO NOT EDIT.
// This file was generated by robots at 2020-04-07 21:49:09.746002 +0000 UTC

package optional

import (
	"encoding/json"
	"errors"
)

// Bool is an optional bool.
type Bool struct {
	val bool
	set bool
}

// MakeBool creates an optional.Bool from a bool.
func MakeBool(v bool) Bool {
	return Bool{val: v, set: true}
}

// Set sets the bool value.
func (b *Bool) Set(v bool) {
	*b = Bool{val: v, set: true}
}

// Unset unsets the bool value.
func (b *Bool) Unset() {
	*b = Bool{}
}

// Present returns whether or not the value is present.
func (b Bool) Present() bool {
	return b.set
}

// Get returns the bool value or an error if not present.
func (b Bool) Get() (bool, error) {
	if !b.Present() {
		var zero bool
		return zero, errors.New("value not present")
	}
	return b.val, nil
}

// MustGet returns the bool value or panics if not present.
func (b Bool) MustGet() bool {
	if !b.Present() {
		panic("value not present")
	}
	return b.val
}

// OrElse returns the bool value or a default value if the value is not present.
func (b Bool) OrElse(v bool) bool {
	if b.Present() {
		return b.val
	}
	return v
}

// If calls the function f with the value if the value is present.
func (b Bool) If(fn func(bool)) {
	if b.Present() {
		fn(b.val)
	}
}

// MarshalJSON implements the json.Marshaler interface.
func (b Bool) MarshalJSON() ([]byte, error) {
	if b.Present() {
		return json.Marshal(b.val)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (b *Bool) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		b.Unset()
		return nil
	}

	var value bool
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	b.Set(value)
	return nil
}
