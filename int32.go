// Code generated by go generate. DO NOT EDIT.
// This file was generated by robots at 2020-04-08 04:22:23.574831 +0000 UTC

package optional

import (
	"encoding/json"
	"errors"
)

// Int32 is an optional int32.
type Int32 struct {
	val int32
	set bool
}

// MakeInt32 creates an optional.Int32 from a int32.
func MakeInt32(v int32) Int32 {
	return Int32{val: v, set: true}
}

// Set sets the int32 value.
func (i *Int32) Set(v int32) {
	*i = MakeInt32(v)
}

// Unset unsets the int32 value.
func (i *Int32) Unset() {
	*i = Int32{}
}

// Present returns whether or not the value is present.
func (i Int32) Present() bool {
	return i.set
}

// Get returns the int32 value or panics if not present.
func (i Int32) Get() int32 {
	if !i.Present() {
		panic("value not present")
	}
	return i.val
}

// GetOr returns the int32 value or a default value if not present.
func (i Int32) GetOr(v int32) int32 {
	if i.Present() {
		return i.val
	}
	return v
}

// GetOrErr returns the int32 value or an error if not present.
func (i Int32) GetOrErr() (int32, error) {
	if !i.Present() {
		var zero int32
		return zero, errors.New("value not present")
	}
	return i.val, nil
}

// If calls the function fn with the value if the value is present.
func (i Int32) If(fn func(int32)) {
	if i.Present() {
		fn(i.val)
	}
}

// Map applies the function fn to the contained value (if any) and returns a new
// option value.
func (i Int32) Map(fn func(int32) int32) Int32 {
	if i.Present() {
		return MakeInt32(fn(i.val))
	}
	return i
}

// And returns an empty Int32 option value if not present, otherwise returns
// optb.
func (i Int32) And(optb Int32) Int32 {
	if i.Present() {
		return optb
	}
	return Int32{}
}

// Or returns the Int32 option value if present, otherwise returns optb.
func (i Int32) Or(optb Int32) Int32 {
	if i.Present() {
		return i
	}
	return optb
}

// MarshalJSON implements the json.Marshaler interface.
func (i Int32) MarshalJSON() ([]byte, error) {
	if i.Present() {
		return json.Marshal(i.val)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (i *Int32) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		i.Unset()
		return nil
	}

	var value int32
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	i.Set(value)
	return nil
}
