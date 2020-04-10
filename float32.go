// Code generated by go generate. DO NOT EDIT.
// This file was generated by robots at 2020-04-10 23:14:18.939169 +0000 UTC

package optional

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
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
	*f = MakeFloat32(v)
}

// Unset unsets the float32 value.
func (f *Float32) Unset() {
	*f = Float32{}
}

// Present returns whether or not the value is present.
func (f Float32) Present() bool {
	return f.set
}

// Get returns the float32 value or panics if not present.
func (f Float32) Get() float32 {
	if !f.Present() {
		panic("value not present")
	}
	return f.val
}

// GetOr returns the float32 value or a default value if not present.
func (f Float32) GetOr(v float32) float32 {
	if f.Present() {
		return f.val
	}
	return v
}

// GetOrBool returns the float32 value or false if not present.
func (f Float32) GetOrBool() (float32, bool) {
	if !f.Present() {
		var zero float32
		return zero, false
	}
	return f.val, true
}

// GetOrErr returns the float32 value or an error if not present.
func (f Float32) GetOrErr() (float32, error) {
	if !f.Present() {
		var zero float32
		return zero, errors.New("value not present")
	}
	return f.val, nil
}

// If calls the function fn with the value if the value is present.
func (f Float32) If(fn func(float32)) {
	if f.Present() {
		fn(f.val)
	}
}

// Map applies the function fn to the contained value (if any) and returns a new
// option value.
func (f Float32) Map(fn func(float32) float32) Float32 {
	if f.Present() {
		return MakeFloat32(fn(f.val))
	}
	return f
}

// And returns an empty Float32 option value if not present, otherwise returns
// optb.
func (f Float32) And(optb Float32) Float32 {
	if f.Present() {
		return optb
	}
	return Float32{}
}

// Or returns the Float32 option value if present, otherwise returns optb.
func (f Float32) Or(optb Float32) Float32 {
	if f.Present() {
		return f
	}
	return optb
}

// Format implements the fmt.Formatter interface.
func (f Float32) Format(fmtS fmt.State, verb rune) {
	if !f.Present() {
		io.WriteString(fmtS, "none")
		return
	}
	var format string
	switch verb {
	case 'v':
		if fmtS.Flag('+') {
			format = "some(%+v)"
		} else {
			format = "some(%v)"
		}
	case 's':
		format = "some(%v)"
	case 'q':
		format = "\"some(%v)\""
	}
	fmt.Fprintf(fmtS, format, f.val)
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
