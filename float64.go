// Code generated by go generate. DO NOT EDIT.
// This file was generated by robots at 2020-04-07 21:41:00.930402 +0000 UTC

package optional

import (
	"encoding/json"
	"errors"
)

// Float64 is an optional float64.
type Float64 struct {
	val float64
	set bool
}

// MakeFloat64 creates an optional.Float64 from a float64.
func MakeFloat64(v float64) Float64 {
	return Float64{val: v, set: true}
}

// Set sets the float64 value.
func (f *Float64) Set(v float64) {
	*f = Float64{val: v, set: true}
}

// Unset unsets the float64 value.
func (f *Float64) Unset() {
	*f = Float64{}
}

// Get returns the float64 value or an error if not present.
func (f Float64) Get() (float64, error) {
	if !f.Present() {
		var zero float64
		return zero, errors.New("value not present")
	}
	return f.val, nil
}

// Present returns whether or not the value is present.
func (f Float64) Present() bool {
	return f.set
}

// OrElse returns the float64 value or a default value if the value is not present.
func (f Float64) OrElse(v float64) float64 {
	if f.Present() {
		return f.val
	}
	return v
}

// If calls the function f with the value if the value is present.
func (f Float64) If(fn func(float64)) {
	if f.Present() {
		fn(f.val)
	}
}

// MarshalJSON implements the json.Marshaler interface.
func (f Float64) MarshalJSON() ([]byte, error) {
	if f.Present() {
		return json.Marshal(f.val)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (f *Float64) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		f.Unset()
		return nil
	}

	var value float64
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	f.Set(value)
	return nil
}
