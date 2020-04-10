package optional

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt_Get_Present(t *testing.T) {
	o := MakeInt(42)

	assert.True(t, o.Present())
	assert.Equal(t, 42, o.Get())
}

func TestInt_Get_NotPresent(t *testing.T) {
	o := Int{}

	assert.False(t, o.Present())
	assert.Panics(t, func() { o.Get() })
}

func TestInt_GetOr_Present(t *testing.T) {
	o := MakeInt(42)

	v := o.GetOr(99)
	assert.True(t, o.Present())
	assert.Equal(t, 42, v)
}

func TestInt_GetOr_NotPresent(t *testing.T) {
	o := Int{}

	v := o.GetOr(99)
	assert.False(t, o.Present())
	assert.Equal(t, 99, v)
}

func TestInt_GetOrBool_Present(t *testing.T) {
	o := MakeInt(42)

	v, ok := o.GetOrBool()
	assert.True(t, o.Present())
	assert.True(t, ok)
	assert.Equal(t, 42, v)
}

func TestInt_GetOrBool_NotPresent(t *testing.T) {
	o := Int{}

	v, ok := o.GetOrBool()
	assert.False(t, o.Present())
	assert.False(t, ok)
	assert.Equal(t, 0, v)
}

func TestInt_GetOrErr_Present(t *testing.T) {
	o := MakeInt(42)

	v, err := o.GetOrErr()
	assert.True(t, o.Present())
	assert.NoError(t, err)
	assert.Equal(t, 42, v)
}

func TestInt_GetOrErr_NotPresent(t *testing.T) {
	o := Int{}

	v, err := o.GetOrErr()
	assert.False(t, o.Present())
	assert.Error(t, err)
	assert.Equal(t, 0, v)
}

func TestInt_If_Present(t *testing.T) {
	o := MakeInt(42)

	canary := false
	o.If(func(v int) {
		canary = true
	})
	assert.True(t, o.Present())
	assert.True(t, canary)
}

func TestInt_If_NotPresent(t *testing.T) {
	o := Int{}

	canary := false
	o.If(func(v int) {
		canary = true
	})
	assert.False(t, o.Present())
	assert.False(t, canary)
}

func TestInt_Map_Present(t *testing.T) {
	o := MakeInt(42)

	v := o.Map(func(i int) int { return 99 })
	assert.True(t, v.Present())
	assert.Equal(t, 99, v.Get())
}

func TestInt_Map_NotPresent(t *testing.T) {
	o := Int{}

	v := o.Map(func(i int) int { return 99 })
	assert.False(t, v.Present())
}

func TestInt_And_Present(t *testing.T) {
	o := MakeInt(42)

	v := o.And(MakeInt(99))
	assert.True(t, v.Present())
	assert.Equal(t, 99, v.Get())

	v2 := o.And(Int{})
	assert.False(t, v2.Present())
}

func TestInt_And_NotPresent(t *testing.T) {
	o := Int{}

	v := o.And(MakeInt(99))
	assert.False(t, v.Present())

	v2 := o.And(Int{})
	assert.False(t, v2.Present())
}

func TestInt_Or_Present(t *testing.T) {
	o := MakeInt(42)

	v := o.Or(MakeInt(99))
	assert.True(t, v.Present())
	assert.Equal(t, 42, v.Get())

	v2 := o.Or(Int{})
	assert.True(t, v2.Present())
	assert.Equal(t, 42, v2.Get())
}

func TestInt_Or_NotPresent(t *testing.T) {
	o := Int{}

	v := o.Or(MakeInt(99))
	assert.True(t, v.Present())
	assert.Equal(t, 99, v.Get())

	v2 := o.Or(Int{})
	assert.False(t, v2.Present())
}

func TestInt_MarshalJSON(t *testing.T) {
	type fields struct {
		WithValue     Int
		WithZeroValue Int
		WithNoValue   Int
		Unused        Int
	}

	var instance = fields{
		WithValue:     MakeInt(42),
		WithZeroValue: MakeInt(0),
		WithNoValue:   Int{},
	}

	out, err := json.Marshal(instance)
	assert.NoError(t, err)
	assert.Equal(t, `{"WithValue":42,"WithZeroValue":0,"WithNoValue":null,"Unused":null}`, string(out))
}

func TestInt_UnmarshalJSON(t *testing.T) {
	type fields struct {
		WithValue     Int
		WithZeroValue Int
		WithNoValue   Int
		Unused        Int
	}

	var jsonString = `{"WithValue":42,"WithZeroValue":0,"WithNoValue":null}`
	instance := fields{}

	err := json.Unmarshal([]byte(jsonString), &instance)
	assert.NoError(t, err)

	assert.True(t, instance.WithZeroValue.Present())
	assert.Equal(t, 42, instance.WithValue.val)

	assert.True(t, instance.WithZeroValue.Present())
	assert.Equal(t, 0, instance.WithZeroValue.val)

	assert.False(t, instance.WithNoValue.Present())
	assert.Equal(t, 0, instance.WithNoValue.val)

	assert.False(t, instance.Unused.Present())
	assert.Equal(t, 0, instance.Unused.val)
}

func TestInt_UnmarshalJSON_Overwritten(t *testing.T) {
	type fields struct {
		WithValue     Int
		WithZeroValue Int
		WithNoValue   Int
		Unused        Int
	}

	var jsonString = `{"WithValue":42,"WithZeroValue":0,"WithNoValue":null}`
	instance := fields{
		WithValue:     MakeInt(1),
		WithZeroValue: MakeInt(2),
		WithNoValue:   MakeInt(3),
		Unused:        MakeInt(4),
	}

	err := json.Unmarshal([]byte(jsonString), &instance)
	assert.NoError(t, err)

	assert.True(t, instance.WithValue.Present())
	assert.Equal(t, 42, instance.WithValue.val)

	assert.True(t, instance.WithZeroValue.Present())
	assert.Equal(t, 0, instance.WithZeroValue.val)

	assert.False(t, instance.WithNoValue.Present())
	assert.Equal(t, 0, instance.WithNoValue.val)

	assert.True(t, instance.Unused.Present())
	assert.Equal(t, 4, instance.Unused.val)
}
