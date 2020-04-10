// Code generated by go generate. DO NOT EDIT.
// This file was generated by robots at 2020-04-10 23:14:17.291315 +0000 UTC

package optional

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// Byte is an optional byte.
type Byte struct {
	val byte
	set bool
}

// MakeByte creates an optional.Byte from a byte.
func MakeByte(v byte) Byte {
	return Byte{val: v, set: true}
}

// Set sets the byte value.
func (b *Byte) Set(v byte) {
	*b = MakeByte(v)
}

// Unset unsets the byte value.
func (b *Byte) Unset() {
	*b = Byte{}
}

// Present returns whether or not the value is present.
func (b Byte) Present() bool {
	return b.set
}

// Get returns the byte value or panics if not present.
func (b Byte) Get() byte {
	if !b.Present() {
		panic("value not present")
	}
	return b.val
}

// GetOr returns the byte value or a default value if not present.
func (b Byte) GetOr(v byte) byte {
	if b.Present() {
		return b.val
	}
	return v
}

// GetOrBool returns the byte value or false if not present.
func (b Byte) GetOrBool() (byte, bool) {
	if !b.Present() {
		var zero byte
		return zero, false
	}
	return b.val, true
}

// GetOrErr returns the byte value or an error if not present.
func (b Byte) GetOrErr() (byte, error) {
	if !b.Present() {
		var zero byte
		return zero, errors.New("value not present")
	}
	return b.val, nil
}

// If calls the function fn with the value if the value is present.
func (b Byte) If(fn func(byte)) {
	if b.Present() {
		fn(b.val)
	}
}

// Map applies the function fn to the contained value (if any) and returns a new
// option value.
func (b Byte) Map(fn func(byte) byte) Byte {
	if b.Present() {
		return MakeByte(fn(b.val))
	}
	return b
}

// And returns an empty Byte option value if not present, otherwise returns
// optb.
func (b Byte) And(optb Byte) Byte {
	if b.Present() {
		return optb
	}
	return Byte{}
}

// Or returns the Byte option value if present, otherwise returns optb.
func (b Byte) Or(optb Byte) Byte {
	if b.Present() {
		return b
	}
	return optb
}

// Format implements the fmt.Formatter interface.
func (b Byte) Format(fmtS fmt.State, verb rune) {
	if !b.Present() {
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
	fmt.Fprintf(fmtS, format, b.val)
}

// MarshalJSON implements the json.Marshaler interface.
func (b Byte) MarshalJSON() ([]byte, error) {
	if b.Present() {
		return json.Marshal(b.val)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (b *Byte) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		b.Unset()
		return nil
	}

	var value byte
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	b.Set(value)
	return nil
}
