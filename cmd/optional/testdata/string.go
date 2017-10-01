// Code generated by go generate
// This file was generated by robots at 2017-10-01 20:33:33.619161789 +0000 UTC

package optional

import "errors"

// String is an optional string
type String struct {
	value *string
}

// OfString creates a optional.String from a string
func OfString(v string) String {
	return String{&v}
}

// Set sets the string value
func (s String) Set(v string) {
	s.value = &v
}

// Get returns the string value or an error if not present
func (s String) Get() (string, error) {
	if !s.Present() {
		return *s.value, errors.New("value not present")
	}
	return *s.value, nil
}

// Present returns whether or not the value is present
func (s String) Present() bool {
	return s.value != nil
}

// OrElse returns the string value or a default value if the value is not present
func (s String) OrElse(v string) string {
	if s.Present() {
		return *s.value
	}
	return v
}

// If calls the function f with the value if the value is present
func (s String) If(fn func(string)) {
	if s.Present() {
		fn(*s.value)
	}
}
