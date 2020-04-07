// Code generated by go generate. DO NOT EDIT.
// This file was generated by robots at 2020-04-07 21:41:00.494095 +0000 UTC

package optional

import (
	"encoding/json"
	"errors"
)

// Float32 is an optional float32.
type Float32 struct {
	val float32
	set bool
}

// MakeFloat32 creates an optional.Float32 from a float32.
func MakeFloat32(v float32) Float32 {
	return Float32{val: v, set: true}
}

// Set sets the float32 value.
func (f *Float32) Set(v float32) {
	*f = Float32{val: v, set: true}
}

// Unset unsets the float32 value.
func (f *Float32) Unset() {
	*f = Float32{}
}

// Get returns the float32 value or an error if not present.
func (f Float32) Get() (float32, error) {
	if !f.Present() {
		var zero float32
		return zero, errors.New("value not present")
	}
	return f.val, nil
}

// Present returns whether or not the value is present.
func (f Float32) Present() bool {
	return f.set
}

// OrElse returns the float32 value or a default value if the value is not present.
func (f Float32) OrElse(v float32) float32 {
	if f.Present() {
		return f.val
	}
	return v
}

// If calls the function f with the value if the value is present.
func (f Float32) If(fn func(float32)) {
	if f.Present() {
		fn(f.val)
	}
}

// MarshalJSON implements the json.Marshaler interface.
func (f Float32) MarshalJSON() ([]byte, error) {
	if f.Present() {
		return json.Marshal(f.val)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (f *Float32) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		f.Unset()
		return nil
	}

	var value float32
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	f.Set(value)
	return nil
}
