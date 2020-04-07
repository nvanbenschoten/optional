// Code generated by go generate. DO NOT EDIT.
// This file was generated by robots at 2020-04-07 21:41:07.09736 +0000 UTC

package optional

import (
	"encoding/json"
	"errors"
)

// Uintptr is an optional uintptr.
type Uintptr struct {
	val uintptr
	set bool
}

// MakeUintptr creates an optional.Uintptr from a uintptr.
func MakeUintptr(v uintptr) Uintptr {
	return Uintptr{val: v, set: true}
}

// Set sets the uintptr value.
func (u *Uintptr) Set(v uintptr) {
	*u = Uintptr{val: v, set: true}
}

// Unset unsets the uintptr value.
func (u *Uintptr) Unset() {
	*u = Uintptr{}
}

// Get returns the uintptr value or an error if not present.
func (u Uintptr) Get() (uintptr, error) {
	if !u.Present() {
		var zero uintptr
		return zero, errors.New("value not present")
	}
	return u.val, nil
}

// Present returns whether or not the value is present.
func (u Uintptr) Present() bool {
	return u.set
}

// OrElse returns the uintptr value or a default value if the value is not present.
func (u Uintptr) OrElse(v uintptr) uintptr {
	if u.Present() {
		return u.val
	}
	return v
}

// If calls the function f with the value if the value is present.
func (u Uintptr) If(fn func(uintptr)) {
	if u.Present() {
		fn(u.val)
	}
}

// MarshalJSON implements the json.Marshaler interface.
func (u Uintptr) MarshalJSON() ([]byte, error) {
	if u.Present() {
		return json.Marshal(u.val)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (u *Uintptr) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		u.Unset()
		return nil
	}

	var value uintptr
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	u.Set(value)
	return nil
}
