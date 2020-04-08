# Optional

[![Build Status](https://travis-ci.org/nvanbenschoten/optional.svg?branch=master)](https://travis-ci.org/nvanbenschoten/optional)
[![Release](https://img.shields.io/github/release/nvanbenschoten/optional.svg?style=flat-square)](https://github.com/nvanbenschoten/optional/releases/latest)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE.md)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/nvanbenschoten/optional)
[![Go Report Card](https://goreportcard.com/badge/github.com/nvanbenschoten/optional?style=flat-square)](https://goreportcard.com/report/github.com/nvanbenschoten/optional)

Optional is a library that provides option types for the primitive Go types.

It can also be used as a tool to generate option type wrappers around your own types.

## Motivation

In Go, variables declared without an explicit initial value are given their zero value. Most of the time this is what you want, but sometimes you want to be able to tell if a variable was set or if it's just a zero value. That's where [option types](https://en.wikipedia.org/wiki/Option_type) come in handy.

This repository is a fork of github.com/markphelps/optional, which also provides option types for Go primitives and custom types. The driving motivation for this fork is performance. Where the original project uses pointer indirection to implement option types, this project inlines values directly into their option wrapper. This inlining avoids memory allocations, reduces garbage collection pressure, and improves cache locality.

## Inspiration

* Rust [std::option](https://doc.rust-lang.org/std/option/)
* Java [Optional](https://docs.oracle.com/javase/8/docs/api/java/util/Optional.html)
* [https://github.com/leighmcculloch/go-optional](https://github.com/leighmcculloch/go-optional)
* [https://github.com/golang/go/issues/7054](https://github.com/golang/go/issues/7054)

## Tool

### Install

`go get -u github.com/nvanbenschoten/optional/cmd/optional`

### Usage

Typically this process would be run using go generate, like this:

```go
//go:generate optional -type=Foo
```

running this command:

```bash
optional -type=Foo
```

in the same directory will create the file optional_foo.go
containing a definition of:

```go
type OptionalFoo struct {
  ...
}
```

The default type is OptionalT or optionalT (depending on if the type is exported)
and output file is optional_t.go. This can be overridden with the -output flag.

## Library

* [bool](bool.go)
* [byte](byte.go)
* [complex128](complex128.go)
* [complex64](complex64.go)
* [float32](float32.go)
* [float64](float64.go)
* [int](int.go)
* [int16](int16.go)
* [int32](int32.go)
* [int64](int64.go)
* [int8](int8.go)
* [rune](rune.go)
* [string](string.go)
* [uint](uint.go)
* [uint16](uint16.go)
* [uint32](uint32.go)
* [uint64](uint64.go)
* [uint8](uint8.go)
* [uintptr](uintptr.go)
* [error](error.go)

### Usage

```go
package main

import (
	"fmt"

	"github.com/nvanbenschoten/optional"
)

func main() {
	s := optional.MakeString("foo")

	value, err := s.GetOrErr()
	if err != nil {
		// handle error!
	} else {
		fmt.Println(value)
	}

	t := optional.String{}
	fmt.Println(t.GetOr("bar"))
}
```

See [example_test.go](example_test.go) and the [documentation](http://godoc.org/github.com/nvanbenschoten/optional) for more usage.

## Marshalling/Unmarshalling JSON

Option types marshal to/from JSON as you would expect:

### Marshalling

```go
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var value = struct {
		Field optional.String `json:"field,omitempty"`
	}{
		Field: optional.MakeString("bar"),
	}

	out, _ := json.Marshal(value)

	fmt.Println(string(out))
	// outputs: {"field":"bar"}
}
```

### Unmarshalling

```go
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var value = &struct {
		Field optional.String `json:"field,omitempty"`
	}{}

	_ = json.Unmarshal([]byte(`{"field":"bar"}`), value)

	value.Field.If(func(s string) {
		fmt.Println(s)
	})
	// outputs: bar
}
```

See [example_test.go](example_test.go) for more examples.

## Test Coverage

As you can see test coverage is a bit lacking for the library. This is simply because testing generated code is not super easy. I'm currently working on improving test coverage for the generated types, but in the meantime checkout [string_test.go](string_test.go) and [int_test.go](int_test.go) for examples.

Also checkout:

* [example_test.go](example_test.go) for example usage.
* [cmd/optional/golden_test.go](cmd/optional/golden_test.go) for [golden file](https://medium.com/soon-london/testing-with-golden-files-in-go-7fccc71c43d3) based testing of the generator itself.

### Golden Files

If changing the API you may need to update the [golden files](https://medium.com/soon-london/testing-with-golden-files-in-go-7fccc71c43d3) for your tests to pass by running:

`go test ./cmd/optional/... -update`.

## Contributing

1. [Fork it](https://github.com/nvanbenschoten/optional/fork)
1. Create your feature branch (`git checkout -b my-new-feature`)
1. Commit your changes (`git commit -am 'Add some feature'`)
1. Push to the branch (`git push origin my-new-feature`)
1. Create a new Pull Request
