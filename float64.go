// Code generated by go generate. DO NOT EDIT.
// This file was generated by robots at 2020-04-10 23:14:19.372873 +0000 UTC

package optional

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
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
	*f = MakeFloat64(v)
}

// Unset unsets the float64 value.
func (f *Float64) Unset() {
	*f = Float64{}
}

// Present returns whether or not the value is present.
func (f Float64) Present() bool {
	return f.set
}

// Get returns the float64 value or panics if not present.
func (f Float64) Get() float64 {
	if !f.Present() {
		panic("value not present")
	}
	return f.val
}

// GetOr returns the float64 value or a default value if not present.
func (f Float64) GetOr(v float64) float64 {
	if f.Present() {
		return f.val
	}
	return v
}

// GetOrBool returns the float64 value or false if not present.
func (f Float64) GetOrBool() (float64, bool) {
	if !f.Present() {
		var zero float64
		return zero, false
	}
	return f.val, true
}

// GetOrErr returns the float64 value or an error if not present.
func (f Float64) GetOrErr() (float64, error) {
	if !f.Present() {
		var zero float64
		return zero, errors.New("value not present")
	}
	return f.val, nil
}

// If calls the function fn with the value if the value is present.
func (f Float64) If(fn func(float64)) {
	if f.Present() {
		fn(f.val)
	}
}

// Map applies the function fn to the contained value (if any) and returns a new
// option value.
func (f Float64) Map(fn func(float64) float64) Float64 {
	if f.Present() {
		return MakeFloat64(fn(f.val))
	}
	return f
}

// And returns an empty Float64 option value if not present, otherwise returns
// optb.
func (f Float64) And(optb Float64) Float64 {
	if f.Present() {
		return optb
	}
	return Float64{}
}

// Or returns the Float64 option value if present, otherwise returns optb.
func (f Float64) Or(optb Float64) Float64 {
	if f.Present() {
		return f
	}
	return optb
}

// Format implements the fmt.Formatter interface.
func (f Float64) Format(fmtS fmt.State, verb rune) {
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
