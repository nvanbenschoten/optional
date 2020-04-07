// Code generated by go generate
// This file was generated by robots at 2020-04-07 21:38:42.208057 +0000 UTC

package optional

import (
	"encoding/json"
	"errors"
)

// Uint32 is an optional uint32.
type Uint32 struct {
	val uint32
	set bool
}

// MakeUint32 creates an optional.Uint32 from a uint32.
func MakeUint32(v uint32) Uint32 {
	return Uint32{val: v, set: true}
}

// Set sets the uint32 value.
func (u *Uint32) Set(v uint32) {
	*u = Uint32{val: v, set: true}
}

// Unset unsets the uint32 value.
func (u *Uint32) Unset() {
	*u = Uint32{}
}

// Get returns the uint32 value or an error if not present.
func (u Uint32) Get() (uint32, error) {
	if !u.Present() {
		var zero uint32
		return zero, errors.New("value not present")
	}
	return u.val, nil
}

// Present returns whether or not the value is present.
func (u Uint32) Present() bool {
	return u.set
}

// OrElse returns the uint32 value or a default value if the value is not present.
func (u Uint32) OrElse(v uint32) uint32 {
	if u.Present() {
		return u.val
	}
	return v
}

// If calls the function f with the value if the value is present.
func (u Uint32) If(fn func(uint32)) {
	if u.Present() {
		fn(u.val)
	}
}

// MarshalJSON implements the json.Marshaler interface.
func (u Uint32) MarshalJSON() ([]byte, error) {
	if u.Present() {
		return json.Marshal(u.val)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (u *Uint32) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		u.Unset()
		return nil
	}

	var value uint32
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	u.Set(value)
	return nil
}
