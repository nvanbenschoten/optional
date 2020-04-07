// Code generated by go generate
// This file was generated by robots at 2020-04-07 21:38:40.393049 +0000 UTC

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

// Get returns the rune value or an error if not present.
func (r Rune) Get() (rune, error) {
	if !r.Present() {
		var zero rune
		return zero, errors.New("value not present")
	}
	return r.val, nil
}

// Present returns whether or not the value is present.
func (r Rune) Present() bool {
	return r.set
}

// OrElse returns the rune value or a default value if the value is not present.
func (r Rune) OrElse(v rune) rune {
	if r.Present() {
		return r.val
	}
	return v
}

// If calls the function f with the value if the value is present.
func (r Rune) If(fn func(rune)) {
	if r.Present() {
		fn(r.val)
	}
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
