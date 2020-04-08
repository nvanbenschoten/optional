package optional

import (
	"errors"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestBool_Size(t *testing.T) {
	o := MakeBool(true)
	assert.Equal(t, 2, int(unsafe.Sizeof(o)))
}

func TestByte_Size(t *testing.T) {
	o := MakeByte('b')
	assert.Equal(t, 2, int(unsafe.Sizeof(o)))
}

func TestComplex128_Size(t *testing.T) {
	o := MakeComplex128(complex(1, 2))
	assert.Equal(t, 24, int(unsafe.Sizeof(o)))
}

func TestComplex64_Size(t *testing.T) {
	o := MakeComplex64(complex(1, 2))
	assert.Equal(t, 12, int(unsafe.Sizeof(o)))
}

func TestError_Size(t *testing.T) {
	o := MakeError(errors.New("error"))
	assert.Equal(t, 24, int(unsafe.Sizeof(o)))
}

func TestFloat32_Size(t *testing.T) {
	o := MakeFloat32(99.9)
	assert.Equal(t, 8, int(unsafe.Sizeof(o)))
}

func TestFloat64_Size(t *testing.T) {
	o := MakeFloat64(99.9)
	assert.Equal(t, 16, int(unsafe.Sizeof(o)))
}

func TestInt_Size(t *testing.T) {
	o := MakeInt(99)
	assert.Equal(t, 16, int(unsafe.Sizeof(o)))
}

func TestInt16_Size(t *testing.T) {
	o := MakeInt16(99)
	assert.Equal(t, 4, int(unsafe.Sizeof(o)))
}

func TestInt32_Size(t *testing.T) {
	o := MakeInt32(99)
	assert.Equal(t, 8, int(unsafe.Sizeof(o)))
}

func TestInt64_Size(t *testing.T) {
	o := MakeInt64(99)
	assert.Equal(t, 16, int(unsafe.Sizeof(o)))
}

func TestInt8_Size(t *testing.T) {
	o := MakeInt8(99)
	assert.Equal(t, 2, int(unsafe.Sizeof(o)))
}

func TestRune_Size(t *testing.T) {
	o := MakeRune('r')
	assert.Equal(t, 8, int(unsafe.Sizeof(o)))
}

func TestString_Size(t *testing.T) {
	o := MakeString("foo")
	assert.Equal(t, 24, int(unsafe.Sizeof(o)))
}

func TestUint_Size(t *testing.T) {
	o := MakeUint(99)
	assert.Equal(t, 16, int(unsafe.Sizeof(o)))
}

func TestUint16_Size(t *testing.T) {
	o := MakeUint16(99)
	assert.Equal(t, 4, int(unsafe.Sizeof(o)))
}

func TestUint32_Size(t *testing.T) {
	o := MakeUint32(99)
	assert.Equal(t, 8, int(unsafe.Sizeof(o)))
}

func TestUint64_Size(t *testing.T) {
	o := MakeUint64(99)
	assert.Equal(t, 16, int(unsafe.Sizeof(o)))
}

func TestUint8_Size(t *testing.T) {
	o := MakeUint8(99)
	assert.Equal(t, 2, int(unsafe.Sizeof(o)))
}

func TestUintptr_Size(t *testing.T) {
	o := MakeUintptr(99)
	assert.Equal(t, 16, int(unsafe.Sizeof(o)))
}
