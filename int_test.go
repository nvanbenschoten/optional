package optional

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt_Get_Present(t *testing.T) {
	o := MakeInt(42)

	v, err := o.Get()
	assert.True(t, o.Present())
	assert.NoError(t, err)
	assert.Equal(t, 42, v)
}

func TestInt_Get_NotPresent(t *testing.T) {
	o := Int{}
	var zero int

	v, err := o.Get()
	assert.False(t, o.Present())
	assert.Error(t, err)
	assert.Equal(t, zero, v)
}

func TestInt_MustGet_Present(t *testing.T) {
	o := MakeInt(42)

	assert.True(t, o.Present())
	assert.Equal(t, 42, o.MustGet())
}

func TestInt_MustGet_NotPresent(t *testing.T) {
	o := Int{}

	assert.False(t, o.Present())
	assert.Panics(t, func() { o.MustGet() })
}

func TestInt_OrElse_Present(t *testing.T) {
	o := MakeInt(42)

	v := o.OrElse(99)
	assert.True(t, o.Present())
	assert.Equal(t, 42, v)
}

func TestInt_OrElse_NotPresent(t *testing.T) {
	o := Int{}

	v := o.OrElse(99)
	assert.False(t, o.Present())
	assert.Equal(t, 99, v)
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
