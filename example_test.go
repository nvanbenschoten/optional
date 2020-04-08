package optional_test

import (
	"encoding/json"
	"fmt"

	"github.com/nvanbenschoten/optional"
)

func Example_present() {
	values := []optional.String{
		optional.MakeString("foo"),
		optional.MakeString(""),
		optional.MakeString("bar"),
		{},
	}

	for _, v := range values {
		fmt.Println(v.Present())
	}
	// Output:
	// true
	// true
	// true
	// false
}

func Example_get() {
	values := []optional.String{
		optional.MakeString("foo"),
		optional.MakeString(""),
		optional.MakeString("bar"),
		{},
	}

	for _, v := range values {
		value := func() (s string) {
			defer func() {
				if r := recover(); r != nil {
					s = "[panics]"
				}
			}()
			return v.Get()
		}()
		fmt.Println(value)
	}
	// Output:
	// foo
	//
	// bar
	// [panics]
}

func Example_getOr() {
	values := []optional.String{
		optional.MakeString("foo"),
		optional.MakeString(""),
		optional.MakeString("bar"),
		{},
	}

	for _, v := range values {
		fmt.Println(v.GetOr("not present"))
	}
	// Output:
	// foo
	//
	// bar
	// not present
}

func Example_getOrErr() {
	values := []optional.String{
		optional.MakeString("foo"),
		optional.MakeString(""),
		optional.MakeString("bar"),
		{},
	}

	for _, v := range values {
		value, err := v.GetOrErr()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(value)
		}
	}
	// Output:
	// foo
	//
	// bar
	// value not present
}

func Example_if() {
	values := []optional.String{
		optional.MakeString("foo"),
		optional.MakeString(""),
		optional.MakeString("bar"),
		{},
	}

	for _, v := range values {
		v.If(func(s string) {
			fmt.Printf("called for %q\n", s)
		})
	}
	// Output:
	// called for "foo"
	// called for ""
	// called for "bar"
}

func Example_map() {
	values := []optional.String{
		optional.MakeString("foo"),
		optional.MakeString(""),
		optional.MakeString("bar"),
		{},
	}

	for _, v := range values {
		v2 := v.Map(func(s string) string {
			return fmt.Sprintf("updated %q", s)
		})

		value, err := v2.GetOrErr()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(value)
		}
	}
	// Output:
	// updated "foo"
	// updated ""
	// updated "bar"
	// value not present
}

func Example_and() {
	values := []optional.String{
		optional.MakeString("foo"),
		optional.MakeString(""),
		optional.MakeString("bar"),
		{},
	}

	for _, v := range values {
		v2 := optional.MakeString("other")
		v3 := v.And(v2)

		value, err := v3.GetOrErr()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(value)
		}
	}
	// Output:
	// other
	// other
	// other
	// value not present
}

func Example_or() {
	values := []optional.String{
		optional.MakeString("foo"),
		optional.MakeString(""),
		optional.MakeString("bar"),
		{},
	}

	for _, v := range values {
		v2 := optional.MakeString("other")
		v3 := v.Or(v2)

		value, err := v3.GetOrErr()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(value)
		}
	}
	// Output:
	// foo
	//
	// bar
	// other
}

func Example_set() {
	var (
		values = []string{
			"foo",
			"",
			"bar",
		}

		s = optional.MakeString("baz")
	)

	for _, v := range values {
		s.Set(v)
		value, err := s.GetOrErr()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(value)
		}
	}
	// Output:
	// foo
	//
	// bar
}

func Example_unset() {
	values := []optional.String{
		optional.MakeString("foo"),
		optional.MakeString(""),
		optional.MakeString("bar"),
		{},
	}

	for _, v := range values {
		v.Unset()
		value, err := v.GetOrErr()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(value)
		}
	}
	// Output:
	// value not present
	// value not present
	// value not present
	// value not present
}

func Example_marshalJSON() {
	type example struct {
		Field    *optional.String `json:"field,omitempty"`
		FieldTwo *optional.String `json:"field_two"`
	}

	var values = []optional.String{
		optional.MakeString("foo"),
		optional.MakeString(""),
		optional.MakeString("bar"),
	}

	for _, v := range values {
		out, _ := json.Marshal(&example{
			Field:    &v,
			FieldTwo: &v,
		})
		fmt.Println(string(out))
	}

	out, _ := json.Marshal(&example{})
	fmt.Println(string(out))
	// Output:
	// {"field":"foo","field_two":"foo"}
	// {"field":"","field_two":""}
	// {"field":"bar","field_two":"bar"}
	// {"field_two":null}
}

func Example_unmarshalJSON() {
	var values = []string{
		`{"field":"foo","field_two":"foo"}`,
		`{"field":"","field_two":""}`,
		`{"field":"null","field_two":"null"}`,
		`{"field":"bar","field_two":"bar"}`,
		"{}",
	}

	for _, v := range values {
		var o = &struct {
			Field    optional.String `json:"field,omitempty"`
			FieldTwo optional.String `json:"field_two"`
		}{}

		if err := json.Unmarshal([]byte(v), o); err != nil {
			fmt.Println(err)
		}

		o.Field.If(func(s string) {
			fmt.Println(s)
		})

		o.FieldTwo.If(func(s string) {
			fmt.Println(s)
		})
	}
	// Output:
	// foo
	// foo
	//

	//
	// bar
	// bar
}
