// Code generated by go generate. DO NOT EDIT.
// This file was generated by robots at 2020-04-20 22:48:08.965284 +0000 UTC

package bar

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// optionalBarPtr is an optional *bar.
type optionalBarPtr struct {
	val *bar
	set bool
}

// MakeoptionalBarPtr creates an optional.optionalBarPtr from a *bar.
func MakeoptionalBarPtr(v *bar) optionalBarPtr {
	return optionalBarPtr{val: v, set: true}
}

// Set sets the *bar value.
func (o *optionalBarPtr) Set(v *bar) {
	*o = MakeoptionalBarPtr(v)
}

// Unset unsets the *bar value.
func (o *optionalBarPtr) Unset() {
	*o = optionalBarPtr{}
}

// Present returns whether or not the value is present.
func (o optionalBarPtr) Present() bool {
	return o.set
}

// Get returns the *bar value or panics if not present.
func (o optionalBarPtr) Get() *bar {
	if !o.Present() {
		panic("value not present")
	}
	return o.val
}

// GetOr returns the *bar value or a default value if not present.
func (o optionalBarPtr) GetOr(v *bar) *bar {
	if o.Present() {
		return o.val
	}
	return v
}

// GetOrBool returns the *bar value or false if not present.
func (o optionalBarPtr) GetOrBool() (*bar, bool) {
	if !o.Present() {
		var zero *bar
		return zero, false
	}
	return o.val, true
}

// GetOrErr returns the *bar value or an error if not present.
func (o optionalBarPtr) GetOrErr() (*bar, error) {
	if !o.Present() {
		var zero *bar
		return zero, errors.New("value not present")
	}
	return o.val, nil
}

// If calls the function fn with the value if the value is present.
func (o optionalBarPtr) If(fn func(*bar)) {
	if o.Present() {
		fn(o.val)
	}
}

// Map applies the function fn to the contained value (if any) and returns a new
// option value.
func (o optionalBarPtr) Map(fn func(*bar) *bar) optionalBarPtr {
	if o.Present() {
		return MakeoptionalBarPtr(fn(o.val))
	}
	return o
}

// And returns an empty optionalBarPtr option value if not present, otherwise returns
// optb.
func (o optionalBarPtr) And(optb optionalBarPtr) optionalBarPtr {
	if o.Present() {
		return optb
	}
	return optionalBarPtr{}
}

// Or returns the optionalBarPtr option value if present, otherwise returns optb.
func (o optionalBarPtr) Or(optb optionalBarPtr) optionalBarPtr {
	if o.Present() {
		return o
	}
	return optb
}

// Format implements the fmt.Formatter interface.
func (o optionalBarPtr) Format(fmtS fmt.State, verb rune) {
	if !o.Present() {
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
	fmt.Fprintf(fmtS, format, o.val)
}

// MarshalJSON implements the json.Marshaler interface.
func (o optionalBarPtr) MarshalJSON() ([]byte, error) {
	if o.Present() {
		return json.Marshal(o.val)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (o *optionalBarPtr) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.Unset()
		return nil
	}

	var value *bar
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	o.Set(value)
	return nil
}
