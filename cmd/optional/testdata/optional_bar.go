// Code generated by go generate. DO NOT EDIT.
// This file was generated by robots at 2020-04-07 21:49:01.432812 +0000 UTC

package bar

import (
	"encoding/json"
	"errors"
)

// optionalBar is an optional bar.
type optionalBar struct {
	val bar
	set bool
}

// MakeoptionalBar creates an optional.optionalBar from a bar.
func MakeoptionalBar(v bar) optionalBar {
	return optionalBar{val: v, set: true}
}

// Set sets the bar value.
func (o *optionalBar) Set(v bar) {
	*o = optionalBar{val: v, set: true}
}

// Unset unsets the bar value.
func (o *optionalBar) Unset() {
	*o = optionalBar{}
}

// Present returns whether or not the value is present.
func (o optionalBar) Present() bool {
	return o.set
}

// Get returns the bar value or an error if not present.
func (o optionalBar) Get() (bar, error) {
	if !o.Present() {
		var zero bar
		return zero, errors.New("value not present")
	}
	return o.val, nil
}

// MustGet returns the bar value or panics if not present.
func (o optionalBar) MustGet() bar {
	if !o.Present() {
		panic("value not present")
	}
	return o.val
}

// OrElse returns the bar value or a default value if the value is not present.
func (o optionalBar) OrElse(v bar) bar {
	if o.Present() {
		return o.val
	}
	return v
}

// If calls the function f with the value if the value is present.
func (o optionalBar) If(fn func(bar)) {
	if o.Present() {
		fn(o.val)
	}
}

// MarshalJSON implements the json.Marshaler interface.
func (o optionalBar) MarshalJSON() ([]byte, error) {
	if o.Present() {
		return json.Marshal(o.val)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (o *optionalBar) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.Unset()
		return nil
	}

	var value bar
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	o.Set(value)
	return nil
}
