// Code generated by go generate
// This file was generated by robots at 2020-04-07 21:28:59.351175 +0000 UTC

package optional

import (
	"encoding/json"
	"errors"
)

// Uint is an optional uint.
type Uint struct {
	val uint
	set bool
}

// MakeUint creates an optional.Uint from a uint.
func MakeUint(v uint) Uint {
	return Uint{val: v, set: true}
}

// Set sets the uint value.
func (u *Uint) Set(v uint) {
	*u = Uint{val: v, set: true}
}

// Unset unsets the uint value.
func (u *Uint) Unset() {
	*u = Uint{}
}

// Get returns the uint value or an error if not present.
func (u Uint) Get() (uint, error) {
	if !u.Present() {
		var zero uint
		return zero, errors.New("value not present")
	}
	return u.val, nil
}

// Present returns whether or not the value is present.
func (u Uint) Present() bool {
	return u.set
}

// OrElse returns the uint value or a default value if the value is not present.
func (u Uint) OrElse(v uint) uint {
	if u.Present() {
		return u.val
	}
	return v
}

// If calls the function f with the value if the value is present.
func (u Uint) If(fn func(uint)) {
	if u.Present() {
		fn(u.val)
	}
}

func (u Uint) MarshalJSON() ([]byte, error) {
	if u.Present() {
		return json.Marshal(u.val)
	}
	return json.Marshal(nil)
}

func (u *Uint) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		u.Unset()
		return nil
	}

	var value uint
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	u.Set(value)
	return nil
}
