// Code generated by go generate. DO NOT EDIT.
// This file was generated by robots at 2020-04-10 23:14:18.107742 +0000 UTC

package optional

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// Complex64 is an optional complex64.
type Complex64 struct {
	val complex64
	set bool
}

// MakeComplex64 creates an optional.Complex64 from a complex64.
func MakeComplex64(v complex64) Complex64 {
	return Complex64{val: v, set: true}
}

// Set sets the complex64 value.
func (c *Complex64) Set(v complex64) {
	*c = MakeComplex64(v)
}

// Unset unsets the complex64 value.
func (c *Complex64) Unset() {
	*c = Complex64{}
}

// Present returns whether or not the value is present.
func (c Complex64) Present() bool {
	return c.set
}

// Get returns the complex64 value or panics if not present.
func (c Complex64) Get() complex64 {
	if !c.Present() {
		panic("value not present")
	}
	return c.val
}

// GetOr returns the complex64 value or a default value if not present.
func (c Complex64) GetOr(v complex64) complex64 {
	if c.Present() {
		return c.val
	}
	return v
}

// GetOrBool returns the complex64 value or false if not present.
func (c Complex64) GetOrBool() (complex64, bool) {
	if !c.Present() {
		var zero complex64
		return zero, false
	}
	return c.val, true
}

// GetOrErr returns the complex64 value or an error if not present.
func (c Complex64) GetOrErr() (complex64, error) {
	if !c.Present() {
		var zero complex64
		return zero, errors.New("value not present")
	}
	return c.val, nil
}

// If calls the function fn with the value if the value is present.
func (c Complex64) If(fn func(complex64)) {
	if c.Present() {
		fn(c.val)
	}
}

// Map applies the function fn to the contained value (if any) and returns a new
// option value.
func (c Complex64) Map(fn func(complex64) complex64) Complex64 {
	if c.Present() {
		return MakeComplex64(fn(c.val))
	}
	return c
}

// And returns an empty Complex64 option value if not present, otherwise returns
// optb.
func (c Complex64) And(optb Complex64) Complex64 {
	if c.Present() {
		return optb
	}
	return Complex64{}
}

// Or returns the Complex64 option value if present, otherwise returns optb.
func (c Complex64) Or(optb Complex64) Complex64 {
	if c.Present() {
		return c
	}
	return optb
}

// Format implements the fmt.Formatter interface.
func (c Complex64) Format(fmtS fmt.State, verb rune) {
	if !c.Present() {
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
	fmt.Fprintf(fmtS, format, c.val)
}

// MarshalJSON implements the json.Marshaler interface.
func (c Complex64) MarshalJSON() ([]byte, error) {
	if c.Present() {
		return json.Marshal(c.val)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (c *Complex64) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		c.Unset()
		return nil
	}

	var value complex64
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	c.Set(value)
	return nil
}
