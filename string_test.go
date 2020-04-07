package optional

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString_Get_Present(t *testing.T) {
	o := MakeString("foo")

	v, err := o.Get()
	assert.True(t, o.Present())
	assert.NoError(t, err)
	assert.Equal(t, "foo", v)
}

func TestString_Get_NotPresent(t *testing.T) {
	o := String{}
	var zero string

	v, err := o.Get()
	assert.False(t, o.Present())
	assert.Error(t, err)
	assert.Equal(t, zero, v)
}

func TestString_MustGet_Present(t *testing.T) {
	o := MakeString("foo")

	assert.True(t, o.Present())
	assert.Equal(t, "foo", o.MustGet())
}

func TestString_MustGet_NotPresent(t *testing.T) {
	o := String{}

	assert.False(t, o.Present())
	assert.Panics(t, func() { o.MustGet() })
}

func TestString_OrElse_Present(t *testing.T) {
	o := MakeString("foo")

	v := o.OrElse("bar")
	assert.True(t, o.Present())
	assert.Equal(t, "foo", v)
}

func TestString_OrElse_NotPresent(t *testing.T) {
	o := String{}

	v := o.OrElse("bar")
	assert.False(t, o.Present())
	assert.Equal(t, "bar", v)
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
