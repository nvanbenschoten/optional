// Code generated by go generate. DO NOT EDIT.
// This file was generated by robots at 2020-04-10 22:06:18.967396 +0000 UTC

package optional

import (
	"encoding/json"
	"errors"
)

// Int is an optional int.
type Int struct {
	val int
	set bool
}

// MakeInt creates an optional.Int from a int.
func MakeInt(v int) Int {
	return Int{val: v, set: true}
}

// Set sets the int value.
func (i *Int) Set(v int) {
	*i = MakeInt(v)
}

// Unset unsets the int value.
func (i *Int) Unset() {
	*i = Int{}
}

// Present returns whether or not the value is present.
func (i Int) Present() bool {
	return i.set
}

// Get returns the int value or panics if not present.
func (i Int) Get() int {
	if !i.Present() {
		panic("value not present")
	}
	return i.val
}

// GetOr returns the int value or a default value if not present.
func (i Int) GetOr(v int) int {
	if i.Present() {
		return i.val
	}
	return v
}

// GetOrBool returns the int value or false if not present.
func (i Int) GetOrBool() (int, bool) {
	if !i.Present() {
		var zero int
		return zero, false
	}
	return i.val, true
}

// GetOrErr returns the int value or an error if not present.
func (i Int) GetOrErr() (int, error) {
	if !i.Present() {
		var zero int
		return zero, errors.New("value not present")
	}
	return i.val, nil
}

// If calls the function fn with the value if the value is present.
func (i Int) If(fn func(int)) {
	if i.Present() {
		fn(i.val)
	}
}

// Map applies the function fn to the contained value (if any) and returns a new
// option value.
func (i Int) Map(fn func(int) int) Int {
	if i.Present() {
		return MakeInt(fn(i.val))
	}
	return i
}

// And returns an empty Int option value if not present, otherwise returns
// optb.
func (i Int) And(optb Int) Int {
	if i.Present() {
		return optb
	}
	return Int{}
}

// Or returns the Int option value if present, otherwise returns optb.
func (i Int) Or(optb Int) Int {
	if i.Present() {
		return i
	}
	return optb
}

// MarshalJSON implements the json.Marshaler interface.
func (i Int) MarshalJSON() ([]byte, error) {
	if i.Present() {
		return json.Marshal(i.val)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (i *Int) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		i.Unset()
		return nil
	}

	var value int
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	i.Set(value)
	return nil
}
