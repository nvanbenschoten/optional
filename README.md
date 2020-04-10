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

## Library

The library provides option types each primitive Go type:

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

	if s.Present() {
		fmt.Println(s.Get())
	} else {
		fmt.Println("missing")
	}

	if val, ok := s.GetOrBool(); ok {
		fmt.Println(val)
	} else {
		fmt.Println("missing")
	}

	if val, err := s.GetOrErr(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val)
	}

	t := optional.String{}
	fmt.Println(t.GetOr("bar"))
}
```

See [example_test.go](example_test.go) and the [documentation](http://godoc.org/github.com/nvanbenschoten/optional) for more usage.

### Memory Layout

Each option type is composed of its value type and a `bool` inlined into a single struct. For example, the option type for an `int32` is represented as `struct{int32, bool}`. The precise memory layout that this structure translates to is subject to Go's [type alignment and padding rules](https://golang.org/ref/spec#Size_and_alignment_guarantees), which make little in the way of guarantees around struct field alignment.

In practice, however, the memory layout for this struct looks like:

```
+---+---+---+---+---+---+---+---+
| i | i | i | i | b | p | p | p |
+---+---+---+---+---+---+---+---+

i = int32   byte
b = bool    byte
p = padding byte

total size = 8 bytes
```

The following table lists the memory footprint in bytes of each of the option types provided by the library when compiled for a 64-bit CPU architecture using Go 1.13.9:

| Type       | Size (bytes) |
|------------|-------------:|
| Bool       |            2 |
| Byte       |            2 |
| Complex128 |           24 |
| Complex64  |           12 |
| Error      |           24 |
| Float32    |            8 |
| Float64    |           16 |
| Int        |           16 |
| Int16      |            4 |
| Int32      |            8 |
| Int64      |           16 |
| Int8       |            2 |
| Rune       |            8 |
| String     |           24 |
| Uint       |           16 |
| Uint16     |            4 |
| Uint32     |            8 |
| Uint64     |           16 |
| Uint8      |            2 |
| Uintptr    |           16 |

These sizes will differ depending on target CPU architecture and may change in future compiler versions. Changes here will break the tests in [size_test.go](size_test.go).

No effort has been made to specialize the implementation of specific option types in order to optimize for memory size. Future work may explore such optimizations in a manner akin to Rust's ["niche-filling strategy"](https://github.com/rust-lang/rust/pull/45225).

## Tool

The tool can generate option type wrappers around your own types.

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
