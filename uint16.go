// Code generated by go generate
// This file was generated by robots at 2020-04-07 21:28:59.778649 +0000 UTC

package optional

import (
	"encoding/json"
	"errors"
)

// Uint16 is an optional uint16.
type Uint16 struct {
	val uint16
	set bool
}

// MakeUint16 creates an optional.Uint16 from a uint16.
func MakeUint16(v uint16) Uint16 {
	return Uint16{val: v, set: true}
}

// Set sets the uint16 value.
func (u *Uint16) Set(v uint16) {
	*u = Uint16{val: v, set: true}
}

// Unset unsets the uint16 value.
func (u *Uint16) Unset() {
	*u = Uint16{}
}

// Get returns the uint16 value or an error if not present.
func (u Uint16) Get() (uint16, error) {
	if !u.Present() {
		var zero uint16
		return zero, errors.New("value not present")
	}
	return u.val, nil
}

// Present returns whether or not the value is present.
func (u Uint16) Present() bool {
	return u.set
}

// OrElse returns the uint16 value or a default value if the value is not present.
func (u Uint16) OrElse(v uint16) uint16 {
	if u.Present() {
		return u.val
	}
	return v
}

// If calls the function f with the value if the value is present.
func (u Uint16) If(fn func(uint16)) {
	if u.Present() {
		fn(u.val)
	}
}

func (u Uint16) MarshalJSON() ([]byte, error) {
	if u.Present() {
		return json.Marshal(u.val)
	}
	return json.Marshal(nil)
}

func (u *Uint16) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		u.Unset()
		return nil
	}

	var value uint16
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	u.Set(value)
	return nil
}
