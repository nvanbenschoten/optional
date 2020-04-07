// Code generated by go generate
// This file was generated by robots at 2020-04-07 21:38:39.331442 +0000 UTC

package optional

import (
	"encoding/json"
	"errors"
)

// Int64 is an optional int64.
type Int64 struct {
	val int64
	set bool
}

// MakeInt64 creates an optional.Int64 from a int64.
func MakeInt64(v int64) Int64 {
	return Int64{val: v, set: true}
}

// Set sets the int64 value.
func (i *Int64) Set(v int64) {
	*i = Int64{val: v, set: true}
}

// Unset unsets the int64 value.
func (i *Int64) Unset() {
	*i = Int64{}
}

// Get returns the int64 value or an error if not present.
func (i Int64) Get() (int64, error) {
	if !i.Present() {
		var zero int64
		return zero, errors.New("value not present")
	}
	return i.val, nil
}

// Present returns whether or not the value is present.
func (i Int64) Present() bool {
	return i.set
}

// OrElse returns the int64 value or a default value if the value is not present.
func (i Int64) OrElse(v int64) int64 {
	if i.Present() {
		return i.val
	}
	return v
}

// If calls the function f with the value if the value is present.
func (i Int64) If(fn func(int64)) {
	if i.Present() {
		fn(i.val)
	}
}

// MarshalJSON implements the json.Marshaler interface.
func (i Int64) MarshalJSON() ([]byte, error) {
	if i.Present() {
		return json.Marshal(i.val)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (i *Int64) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		i.Unset()
		return nil
	}

	var value int64
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	i.Set(value)
	return nil
}
