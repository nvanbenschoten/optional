// Code generated by go generate. DO NOT EDIT.
// This file was generated by robots at 2020-04-08 04:05:17.724566 +0000 UTC

package optional

import (
	"encoding/json"
	"errors"
)

// Rune is an optional rune.
type Rune struct {
	val rune
	set bool
}

// MakeRune creates an optional.Rune from a rune.
func MakeRune(v rune) Rune {
	return Rune{val: v, set: true}
}

// Set sets the rune value.
func (r *Rune) Set(v rune) {
	*r = Rune{val: v, set: true}
}

// Unset unsets the rune value.
func (r *Rune) Unset() {
	*r = Rune{}
}

// Present returns whether or not the value is present.
func (r Rune) Present() bool {
	return r.set
}

// Get returns the rune value or panics if not present.
func (r Rune) Get() rune {
	if !r.Present() {
		panic("value not present")
	}
	return r.val
}

// GetOr returns the rune value or a default value if not present.
func (r Rune) GetOr(v rune) rune {
	if r.Present() {
		return r.val
	}
	return v
}

// GetOrErr returns the rune value or an error if not present.
func (r Rune) GetOrErr() (rune, error) {
	if !r.Present() {
		var zero rune
		return zero, errors.New("value not present")
	}
	return r.val, nil
}

// If calls the function fn with the value if the value is present.
func (r Rune) If(fn func(rune)) {
	if r.Present() {
		fn(r.val)
	}
}

// Map applies the function fn to the contained value (if any) and returns a new
// option value.
func (r Rune) Map(fn func(rune) rune) Rune {
	if r.Present() {
		return MakeRune(fn(r.val))
	}
	return r
}

// And returns an empty Rune option value if not present, otherwise returns
// optb.
func (r Rune) And(optb Rune) Rune {
	if r.Present() {
		return optb
	}
	return Rune{}
}

// Or returns the Rune option value if present, otherwise returns optb.
func (r Rune) Or(optb Rune) Rune {
	if r.Present() {
		return r
	}
	return optb
}

// MarshalJSON implements the json.Marshaler interface.
func (r Rune) MarshalJSON() ([]byte, error) {
	if r.Present() {
		return json.Marshal(r.val)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (r *Rune) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		r.Unset()
		return nil
	}

	var value rune
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	r.Set(value)
	return nil
}
