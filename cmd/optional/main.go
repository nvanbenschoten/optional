// Optional is a tool that generates 'optional' type wrappers around a given type T.
//
// Typically this process would be run using go generate, like this:
//
//	//go:generate optional -type=Foo
//
// running this command
//
//	optional -type=Foo
//
// in the same directory will create the file optional_foo.go
// containing a definition of
//
//	type OptionalFoo struct {
//		...
//	}
//
// The default type is OptionalT or optionalT (depending on if the type is exported)
// and output file is optional_t.go. This can be overridden with the -output flag.
//
package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/build"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
	"time"
)

type generator struct {
	packageName string
	outputName  string
	typeName    string
}

func (g *generator) generate() ([]byte, error) {

	var (
		t = template.Must(template.New("").Parse(tmpl))

		data = struct {
			Timestamp    time.Time
			PackageName  string
			TypeName     string
			OutputName   string
			VariableName string
		}{
			time.Now().UTC(),
			g.packageName,
			g.typeName,
			g.outputName,
			strings.ToLower(string(g.outputName[0])),
		}

		buf bytes.Buffer
		err = t.Execute(&buf, data)
	)

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("optional: ")

	var (
		typeName   = flag.String("type", "", "type name; must be set")
		outputName = flag.String("output", "", "output type and file name; default [o|O]ptional<type> and srcdir/optional_<type>.go")
	)

	flag.Parse()

	if len(*typeName) == 0 {
		flag.Usage()
		os.Exit(2)
	}

	pkg, err := build.Default.ImportDir(".", 0)
	if err != nil {
		log.Fatal(err)
	}

	var (
		filename string
		g        generator
	)

	g.typeName = *typeName
	g.packageName = pkg.Name

	if len(*outputName) == 0 {
		// no output specified, use default optional_<type>

		// TODO: may not be the most reliable method
		exported := strings.Title(g.typeName) == g.typeName

		if exported {
			g.outputName = "Optional" + strings.Title(g.typeName)
		} else {
			g.outputName = "optional" + strings.Title(g.typeName)
		}
		filename = fmt.Sprintf("optional_%s.go", strings.ToLower(g.typeName))
	} else {
		g.outputName = *outputName
		filename = strings.ToLower(g.outputName + ".go")
	}

	src, err := g.generate()
	if err != nil {
		log.Fatal(err)
	}

	if err = ioutil.WriteFile(filename, src, 0644); err != nil {
		log.Fatalf("writing output: %s", err)
	}
}

const tmpl = `// Code generated by go generate. DO NOT EDIT.
// This file was generated by robots at {{ .Timestamp }}

package {{ .PackageName }}

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// {{ .OutputName }} is an optional {{ .TypeName }}.
type {{ .OutputName }} struct {
	val {{ .TypeName }}
	set bool
}

// Make{{ .OutputName }} creates an optional.{{ .OutputName }} from a {{ .TypeName }}.
func Make{{ .OutputName }}(v {{ .TypeName }}) {{ .OutputName }} {
	return {{ .OutputName }}{val: v, set: true}
}

// Set sets the {{ .TypeName }} value.
func ({{ .VariableName }} *{{ .OutputName }}) Set(v {{ .TypeName }}) {
	*{{ .VariableName }} = Make{{ .OutputName }}(v)
}

// Unset unsets the {{ .TypeName }} value.
func ({{ .VariableName }} *{{ .OutputName }}) Unset() {
	*{{ .VariableName }} = {{ .OutputName }}{}
}

// Present returns whether or not the value is present.
func ({{ .VariableName }} {{ .OutputName }}) Present() bool {
	return {{ .VariableName }}.set
}

// Get returns the {{ .TypeName }} value or panics if not present.
func ({{ .VariableName }} {{ .OutputName }}) Get() {{ .TypeName }} {
	if !{{ .VariableName }}.Present() {
		panic("value not present")
	}
	return {{ .VariableName }}.val
}

// GetOr returns the {{ .TypeName }} value or a default value if not present.
func ({{ .VariableName }} {{ .OutputName }}) GetOr(v {{ .TypeName }}) {{ .TypeName }} {
	if {{ .VariableName }}.Present() {
		return {{ .VariableName }}.val
	}
	return v
}

// GetOrBool returns the {{ .TypeName }} value or false if not present.
func ({{ .VariableName }} {{ .OutputName }}) GetOrBool() ({{ .TypeName }}, bool) {
	if !{{ .VariableName }}.Present() {
		var zero {{ .TypeName }}
		return zero, false
	}
	return {{ .VariableName }}.val, true
}

// GetOrErr returns the {{ .TypeName }} value or an error if not present.
func ({{ .VariableName }} {{ .OutputName }}) GetOrErr() ({{ .TypeName }}, error) {
	if !{{ .VariableName }}.Present() {
		var zero {{ .TypeName }}
		return zero, errors.New("value not present")
	}
	return {{ .VariableName }}.val, nil
}

// If calls the function fn with the value if the value is present.
func ({{ .VariableName }} {{ .OutputName }}) If(fn func({{ .TypeName }})) {
	if {{ .VariableName }}.Present() {
		fn({{ .VariableName }}.val)
	}
}

// Map applies the function fn to the contained value (if any) and returns a new
// option value.
func ({{ .VariableName }} {{ .OutputName }}) Map(fn func({{ .TypeName }}) {{ .TypeName }}) {{ .OutputName }} {
	if {{ .VariableName }}.Present() {
		return Make{{ .OutputName }}(fn({{ .VariableName }}.val))
	}
	return {{ .VariableName }}
}

// And returns an empty {{ .OutputName }} option value if not present, otherwise returns
// optb.
func ({{ .VariableName }} {{ .OutputName }}) And(optb {{ .OutputName }}) {{ .OutputName }} {
	if {{ .VariableName }}.Present() {
		return optb
	}
	return {{ .OutputName }}{}
}

// Or returns the {{ .OutputName }} option value if present, otherwise returns optb.
func ({{ .VariableName }} {{ .OutputName }}) Or(optb {{ .OutputName }}) {{ .OutputName }} {
	if {{ .VariableName }}.Present() {
		return {{ .VariableName }}
	}
	return optb
}

// Format implements the fmt.Formatter interface.
func ({{ .VariableName }} {{ .OutputName }}) Format(fmtS fmt.State, verb rune) {
	if !{{ .VariableName }}.Present() {
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
	fmt.Fprintf(fmtS, format, {{ .VariableName }}.val)
}

// MarshalJSON implements the json.Marshaler interface.
func ({{ .VariableName }} {{ .OutputName }}) MarshalJSON() ([]byte, error) {
	if {{ .VariableName }}.Present() {
		return json.Marshal({{ .VariableName }}.val)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func ({{ .VariableName }} *{{ .OutputName }}) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		{{ .VariableName }}.Unset()
		return nil
	}

	var value {{ .TypeName }}
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	{{ .VariableName }}.Set(value)
	return nil
}
`
