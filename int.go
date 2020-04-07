// Code generated by go generate
// This file was generated by robots at 2020-04-07 21:38:37.829232 +0000 UTC

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
	*i = Int{val: v, set: true}
}

// Unset unsets the int value.
func (i *Int) Unset() {
	*i = Int{}
}

// Get returns the int value or an error if not present.
func (i Int) Get() (int, error) {
	if !i.Present() {
		var zero int
		return zero, errors.New("value not present")
	}
	return i.val, nil
}

// Present returns whether or not the value is present.
func (i Int) Present() bool {
	return i.set
}

// OrElse returns the int value or a default value if the value is not present.
func (i Int) OrElse(v int) int {
	if i.Present() {
		return i.val
	}
	return v
}

// If calls the function f with the value if the value is present.
func (i Int) If(fn func(int)) {
	if i.Present() {
		fn(i.val)
	}
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
