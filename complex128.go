// Code generated by go generate. DO NOT EDIT.
// This file was generated by robots at 2020-04-08 04:05:12.831056 +0000 UTC

package optional

import (
	"encoding/json"
	"errors"
)

// Complex128 is an optional complex128.
type Complex128 struct {
	val complex128
	set bool
}

// MakeComplex128 creates an optional.Complex128 from a complex128.
func MakeComplex128(v complex128) Complex128 {
	return Complex128{val: v, set: true}
}

// Set sets the complex128 value.
func (c *Complex128) Set(v complex128) {
	*c = Complex128{val: v, set: true}
}

// Unset unsets the complex128 value.
func (c *Complex128) Unset() {
	*c = Complex128{}
}

// Present returns whether or not the value is present.
func (c Complex128) Present() bool {
	return c.set
}

// Get returns the complex128 value or panics if not present.
func (c Complex128) Get() complex128 {
	if !c.Present() {
		panic("value not present")
	}
	return c.val
}

// GetOr returns the complex128 value or a default value if not present.
func (c Complex128) GetOr(v complex128) complex128 {
	if c.Present() {
		return c.val
	}
	return v
}

// GetOrErr returns the complex128 value or an error if not present.
func (c Complex128) GetOrErr() (complex128, error) {
	if !c.Present() {
		var zero complex128
		return zero, errors.New("value not present")
	}
	return c.val, nil
}

// If calls the function fn with the value if the value is present.
func (c Complex128) If(fn func(complex128)) {
	if c.Present() {
		fn(c.val)
	}
}

// Map applies the function fn to the contained value (if any) and returns a new
// option value.
func (c Complex128) Map(fn func(complex128) complex128) Complex128 {
	if c.Present() {
		return MakeComplex128(fn(c.val))
	}
	return c
}

// And returns an empty Complex128 option value if not present, otherwise returns
// optb.
func (c Complex128) And(optb Complex128) Complex128 {
	if c.Present() {
		return optb
	}
	return Complex128{}
}

// Or returns the Complex128 option value if present, otherwise returns optb.
func (c Complex128) Or(optb Complex128) Complex128 {
	if c.Present() {
		return c
	}
	return optb
}

// MarshalJSON implements the json.Marshaler interface.
func (c Complex128) MarshalJSON() ([]byte, error) {
	if c.Present() {
		return json.Marshal(c.val)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (c *Complex128) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		c.Unset()
		return nil
	}

	var value complex128
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	c.Set(value)
	return nil
}
