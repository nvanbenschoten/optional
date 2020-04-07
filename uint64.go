// Code generated by go generate. DO NOT EDIT.
// This file was generated by robots at 2020-04-07 21:41:06.198304 +0000 UTC

package optional

import (
	"encoding/json"
	"errors"
)

// Uint64 is an optional uint64.
type Uint64 struct {
	val uint64
	set bool
}

// MakeUint64 creates an optional.Uint64 from a uint64.
func MakeUint64(v uint64) Uint64 {
	return Uint64{val: v, set: true}
}

// Set sets the uint64 value.
func (u *Uint64) Set(v uint64) {
	*u = Uint64{val: v, set: true}
}

// Unset unsets the uint64 value.
func (u *Uint64) Unset() {
	*u = Uint64{}
}

// Get returns the uint64 value or an error if not present.
func (u Uint64) Get() (uint64, error) {
	if !u.Present() {
		var zero uint64
		return zero, errors.New("value not present")
	}
	return u.val, nil
}

// Present returns whether or not the value is present.
func (u Uint64) Present() bool {
	return u.set
}

// OrElse returns the uint64 value or a default value if the value is not present.
func (u Uint64) OrElse(v uint64) uint64 {
	if u.Present() {
		return u.val
	}
	return v
}

// If calls the function f with the value if the value is present.
func (u Uint64) If(fn func(uint64)) {
	if u.Present() {
		fn(u.val)
	}
}

// MarshalJSON implements the json.Marshaler interface.
func (u Uint64) MarshalJSON() ([]byte, error) {
	if u.Present() {
		return json.Marshal(u.val)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (u *Uint64) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		u.Unset()
		return nil
	}

	var value uint64
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	u.Set(value)
	return nil
}
