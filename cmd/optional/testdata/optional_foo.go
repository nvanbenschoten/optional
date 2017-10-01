// Code generated by go generate
// This file was generated by robots at 2017-10-01 20:33:33.620123711 +0000 UTC

package foo

import "errors"

// OptionalFoo is an optional Foo
type OptionalFoo struct {
	value *Foo
}

// OfOptionalFoo creates a optional.OptionalFoo from a Foo
func OfOptionalFoo(v Foo) OptionalFoo {
	return OptionalFoo{&v}
}

// Set sets the Foo value
func (o OptionalFoo) Set(v Foo) {
	o.value = &v
}

// Get returns the Foo value or an error if not present
func (o OptionalFoo) Get() (Foo, error) {
	if !o.Present() {
		return *o.value, errors.New("value not present")
	}
	return *o.value, nil
}

// Present returns whether or not the value is present
func (o OptionalFoo) Present() bool {
	return o.value != nil
}

// OrElse returns the Foo value or a default value if the value is not present
func (o OptionalFoo) OrElse(v Foo) Foo {
	if o.Present() {
		return *o.value
	}
	return v
}

// If calls the function f with the value if the value is present
func (o OptionalFoo) If(fn func(Foo)) {
	if o.Present() {
		fn(*o.value)
	}
}
