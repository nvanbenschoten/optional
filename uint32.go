// Code generated by go generate. DO NOT EDIT.
// This file was generated by robots at 2020-04-08 04:05:19.453147 +0000 UTC

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

// Present returns whether or not the value is present.
func (u Uint32) Present() bool {
	return u.set
}

// Get returns the uint32 value or panics if not present.
func (u Uint32) Get() uint32 {
	if !u.Present() {
		panic("value not present")
	}
	return u.val
}

// GetOr returns the uint32 value or a default value if not present.
func (u Uint32) GetOr(v uint32) uint32 {
	if u.Present() {
		return u.val
	}
	return v
}

// GetOrErr returns the uint32 value or an error if not present.
func (u Uint32) GetOrErr() (uint32, error) {
	if !u.Present() {
		var zero uint32
		return zero, errors.New("value not present")
	}
	return u.val, nil
}

// If calls the function fn with the value if the value is present.
func (u Uint32) If(fn func(uint32)) {
	if u.Present() {
		fn(u.val)
	}
}

// Map applies the function fn to the contained value (if any) and returns a new
// option value.
func (u Uint32) Map(fn func(uint32) uint32) Uint32 {
	if u.Present() {
		return MakeUint32(fn(u.val))
	}
	return u
}

// And returns an empty Uint32 option value if not present, otherwise returns
// optb.
func (u Uint32) And(optb Uint32) Uint32 {
	if u.Present() {
		return optb
	}
	return Uint32{}
}

// Or returns the Uint32 option value if present, otherwise returns optb.
func (u Uint32) Or(optb Uint32) Uint32 {
	if u.Present() {
		return u
	}
	return optb
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
