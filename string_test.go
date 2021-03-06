package optional

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString_Get_Present(t *testing.T) {
	o := MakeString("foo")

	assert.True(t, o.Present())
	assert.Equal(t, "foo", o.Get())
}

func TestString_Get_NotPresent(t *testing.T) {
	o := String{}

	assert.False(t, o.Present())
	assert.Panics(t, func() { o.Get() })
}

func TestString_GetOr_Present(t *testing.T) {
	o := MakeString("foo")

	v := o.GetOr("bar")
	assert.True(t, o.Present())
	assert.Equal(t, "foo", v)
}

func TestString_GetOr_NotPresent(t *testing.T) {
	o := String{}

	v := o.GetOr("bar")
	assert.False(t, o.Present())
	assert.Equal(t, "bar", v)
}

func TestString_GetOrBool_Present(t *testing.T) {
	o := MakeString("foo")

	v, ok := o.GetOrBool()
	assert.True(t, o.Present())
	assert.True(t, ok)
	assert.Equal(t, "foo", v)
}

func TestString_GetOrBool_NotPresent(t *testing.T) {
	o := String{}

	v, ok := o.GetOrBool()
	assert.False(t, o.Present())
	assert.False(t, ok)
	assert.Equal(t, "", v)
}

func TestString_GetOrErr_Present(t *testing.T) {
	o := MakeString("foo")

	v, err := o.GetOrErr()
	assert.True(t, o.Present())
	assert.NoError(t, err)
	assert.Equal(t, "foo", v)
}

func TestString_GetOrErr_NotPresent(t *testing.T) {
	o := String{}

	v, err := o.GetOrErr()
	assert.False(t, o.Present())
	assert.Error(t, err)
	assert.Equal(t, "", v)
}

func TestString_If_Present(t *testing.T) {
	o := MakeString("foo")

	canary := false
	o.If(func(v string) {
		canary = true
	})
	assert.True(t, o.Present())
	assert.True(t, canary)
}

func TestString_If_NotPresent(t *testing.T) {
	o := String{}

	canary := false
	o.If(func(v string) {
		canary = true
	})
	assert.False(t, o.Present())
	assert.False(t, canary)
}

func TestString_Map_Present(t *testing.T) {
	o := MakeString("foo")

	v := o.Map(func(s string) string { return "bar" })
	assert.True(t, v.Present())
	assert.Equal(t, "bar", v.Get())
}

func TestString_Map_NotPresent(t *testing.T) {
	o := String{}

	v := o.Map(func(s string) string { return "bar" })
	assert.False(t, v.Present())
}

func TestString_And_Present(t *testing.T) {
	o := MakeString("foo")

	v := o.And(MakeString("bar"))
	assert.True(t, v.Present())
	assert.Equal(t, "bar", v.Get())

	v2 := o.And(String{})
	assert.False(t, v2.Present())
}

func TestString_And_NotPresent(t *testing.T) {
	o := String{}

	v := o.And(MakeString("bar"))
	assert.False(t, v.Present())

	v2 := o.And(String{})
	assert.False(t, v2.Present())
}

func TestString_Or_Present(t *testing.T) {
	o := MakeString("foo")

	v := o.Or(MakeString("bar"))
	assert.True(t, v.Present())
	assert.Equal(t, "foo", v.Get())

	v2 := o.Or(String{})
	assert.True(t, v2.Present())
	assert.Equal(t, "foo", v2.Get())
}

func TestString_Or_NotPresent(t *testing.T) {
	o := String{}

	v := o.Or(MakeString("bar"))
	assert.True(t, v.Present())
	assert.Equal(t, "bar", v.Get())

	v2 := o.Or(String{})
	assert.False(t, v2.Present())
}

func TestString_Format(t *testing.T) {
	type fields struct {
		WithValue     String
		WithZeroValue String
		WithNoValue   String
		Unused        String
	}

	var instance = fields{
		WithValue:     MakeString("foo"),
		WithZeroValue: MakeString(""),
		WithNoValue:   String{},
	}

	out := fmt.Sprintf("%s", instance)
	assert.Equal(t, `{some(foo) some() none none}`, out)
}

func TestString_MarshalJSON(t *testing.T) {
	type fields struct {
		WithValue     String
		WithZeroValue String
		WithNoValue   String
		Unused        String
	}

	var instance = fields{
		WithValue:     MakeString("foo"),
		WithZeroValue: MakeString(""),
		WithNoValue:   String{},
	}

	out, err := json.Marshal(instance)
	assert.NoError(t, err)
	assert.Equal(t, `{"WithValue":"foo","WithZeroValue":"","WithNoValue":null,"Unused":null}`, string(out))
}

func TestString_UnmarshalJSON(t *testing.T) {
	type fields struct {
		WithValue     String
		WithZeroValue String
		WithNoValue   String
		Unused        String
	}

	var jsonString = `{"WithValue":"foo","WithZeroValue":"","WithNoValue":null}`
	instance := fields{}

	err := json.Unmarshal([]byte(jsonString), &instance)
	assert.NoError(t, err)

	assert.True(t, instance.WithZeroValue.Present())
	assert.Equal(t, "foo", instance.WithValue.val)

	assert.True(t, instance.WithZeroValue.Present())
	assert.Equal(t, "", instance.WithZeroValue.val)

	assert.False(t, instance.WithNoValue.Present())
	assert.Equal(t, "", instance.WithNoValue.val)

	assert.False(t, instance.Unused.Present())
	assert.Equal(t, "", instance.Unused.val)
}

func TestString_UnmarshalJSON_Overwritten(t *testing.T) {
	type fields struct {
		WithValue     String
		WithZeroValue String
		WithNoValue   String
		Unused        String
	}

	var jsonString = `{"WithValue":"foo","WithZeroValue":"","WithNoValue":null}`
	instance := fields{
		WithValue:     MakeString("seed_a"),
		WithZeroValue: MakeString("seed_b"),
		WithNoValue:   MakeString("seed_c"),
		Unused:        MakeString("seed_d"),
	}

	err := json.Unmarshal([]byte(jsonString), &instance)
	assert.NoError(t, err)

	assert.True(t, instance.WithValue.Present())
	assert.Equal(t, "foo", instance.WithValue.val)

	assert.True(t, instance.WithZeroValue.Present())
	assert.Equal(t, "", instance.WithZeroValue.val)

	assert.False(t, instance.WithNoValue.Present())
	assert.Equal(t, "", instance.WithNoValue.val)

	assert.True(t, instance.Unused.Present())
	assert.Equal(t, "seed_d", instance.Unused.val)
}
