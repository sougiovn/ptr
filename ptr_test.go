package ptr

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type teststruct struct {
	foo string
	bar int
}

func Test_Ptr(t *testing.T) {
	t.Run("struct", func(t *testing.T) {
		value := teststruct{foo: "baz", bar: 42}

		pointer := Ptr(value)
		assert.Equal(t, value, *pointer)
	})

	t.Run("slice/string", func(t *testing.T) {
		value := []string{"foo", "bar"}

		pointer := Ptr(value)
		assert.Equal(t, value, *pointer)
	})

	t.Run("slice/struct", func(t *testing.T) {
		value := []teststruct{
			{"foo", 42},
			{"bar", 69},
		}

		pointer := Ptr(value)
		assert.Equal(t, value, *pointer)
	})
}

func Test_PtrSlice(t *testing.T) {
	t.Run("string[]", func(t *testing.T) {
		value := []string{"foo", "bar", "baz"}

		pointer := PtrSlice(value)
		require.Len(t, pointer, len(value))
		for i := range value {
			assert.Equal(t, value[i], *pointer[i])
		}
	})

	t.Run("struct[]", func(t *testing.T) {
		value := []teststruct{
			{"foo", 42},
			{"bar", 69},
			{"baz", 99},
		}

		pointer := PtrSlice(value)
		require.Len(t, pointer, len(value))
		for i := range value {
			assert.Equal(t, value[i], *pointer[i])
		}
	})
}

func Test_PtrMap(t *testing.T) {
	t.Run("[string]string", func(t *testing.T) {
		value := map[string]string{
			"foo": "foo",
			"bar": "bar",
			"baz": "baz",
		}

		pointer := PtrMap(value)
		require.Len(t, pointer, len(value))
		for k := range value {
			assert.Equal(t, value[k], *pointer[k])
		}
	})

	t.Run("[int]string", func(t *testing.T) {
		value := map[int]string{
			42: "foo",
			69: "bar",
			99: "baz",
		}

		pointer := PtrMap(value)
		require.Len(t, pointer, len(value))
		for k := range value {
			assert.Equal(t, value[k], *pointer[k])
		}
	})

	t.Run("[bool]int", func(t *testing.T) {
		value := map[bool]int{
			true:  42,
			false: 69,
		}

		pointer := PtrMap(value)
		require.Len(t, pointer, len(value))
		for k := range value {
			assert.Equal(t, value[k], *pointer[k])
		}
	})

	t.Run("[string]struct", func(t *testing.T) {
		value := map[string]teststruct{
			"foo": {"foo", 42},
			"bar": {"bar", 69},
			"baz": {"baz", 420},
		}

		pointer := PtrMap(value)
		require.Len(t, pointer, len(value))
		for k := range value {
			assert.Equal(t, value[k], *pointer[k])
		}
	})
}

func Test_Value(t *testing.T) {
	t.Run("struct", func(t *testing.T) {
		pointer := &teststruct{foo: "baz", bar: 42}

		valueue := Value(pointer)
		assert.Equal(t, *pointer, valueue)
		assert.Equal(t, pointer, &valueue)
	})

	t.Run("slice/string", func(t *testing.T) {
		pointer := &[]string{"foo", "bar"}

		valueue := Value(pointer)
		assert.Equal(t, *pointer, valueue)
		assert.Equal(t, pointer, &valueue)
	})

	t.Run("slice/struct", func(t *testing.T) {
		pointer := &[]teststruct{
			{"foo", 42},
			{"bar", 69},
		}

		valueue := Value(pointer)
		assert.Equal(t, *pointer, valueue)
		assert.Equal(t, pointer, &valueue)
	})

	t.Run("nil", func(t *testing.T) {
		var pointer *teststruct

		valueue := Value(pointer)
		assert.Equal(t, teststruct{}, valueue)
	})
}

func Test_ValueSlice(t *testing.T) {
	t.Run("string[]", func(t *testing.T) {
		foo := "foo"
		bar := "bar"
		baz := "baz"
		pointer := []*string{&foo, &bar, &baz}

		value := ValueSlice(pointer)
		require.Len(t, value, len(pointer))
		for i := range pointer {
			assert.Equal(t, *pointer[i], value[i])
		}
	})

	t.Run("struct[]", func(t *testing.T) {
		foo := teststruct{"foo", 42}
		bar := teststruct{"bar", 69}
		baz := teststruct{"baz", 99}
		pointer := []*teststruct{&foo, &bar, &baz}

		value := ValueSlice(pointer)
		require.Len(t, value, len(pointer))
		for i := range pointer {
			assert.Equal(t, *pointer[i], value[i])
		}
	})
}

func Test_ValueMap(t *testing.T) {
	t.Run("[string]string", func(t *testing.T) {
		foo := "foo"
		bar := "bar"
		baz := "baz"
		pointer := map[string]*string{
			"foo": &foo,
			"bar": &bar,
			"baz": &baz,
		}

		value := ValueMap(pointer)
		require.Len(t, value, len(value))
		for k := range pointer {
			assert.Equal(t, *pointer[k], value[k])
		}
	})

	t.Run("[int]string", func(t *testing.T) {
		foo := "foo"
		bar := "bar"
		baz := "baz"
		pointer := map[int]*string{
			42: &foo,
			69: &bar,
			99: &baz,
		}

		value := ValueMap(pointer)
		require.Len(t, value, len(value))
		for k := range pointer {
			assert.Equal(t, *pointer[k], value[k])
		}
	})

	t.Run("[bool]int", func(t *testing.T) {
		b1 := 42
		b2 := 69
		pointer := map[bool]*int{
			true:  &b1,
			false: &b2,
		}

		value := ValueMap(pointer)
		require.Len(t, value, len(value))
		for k := range pointer {
			assert.Equal(t, *pointer[k], value[k])
		}
	})

	t.Run("[string]struct", func(t *testing.T) {
		foo := teststruct{"foo", 42}
		bar := teststruct{"bar", 69}
		baz := teststruct{"baz", 99}
		pointer := map[string]*teststruct{
			"foo": &foo,
			"bar": &bar,
			"baz": &baz,
		}

		value := ValueMap(pointer)
		require.Len(t, value, len(value))
		for k := range pointer {
			assert.Equal(t, *pointer[k], value[k])
		}
	})
}

// type wrapper code generated by ./generated/main.go
func Test_Byte(t *testing.T) {
	t.Run("byte", func(t *testing.T) {
		value := byte(42)

		pointer := Byte(value)
		assert.Equal(t, value, *pointer)
	})

	t.Run("byte/slice", func(t *testing.T) {
		value := []byte{byte(42), byte(69), byte(99)}

		pointer := ByteSlice(value)
		require.Len(t, pointer, len(value))
		for i := range value {
			assert.Equal(t, value[i], *pointer[i])
		}
	})

	t.Run("byte/map", func(t *testing.T) {
		value := map[string]byte{
			"foo": byte(42), 
			"bar": byte(69), 
			"baz": byte(99),
		}

		pointer := ByteMap(value)
		require.Len(t, pointer, len(value))
		for i := range value {
			assert.Equal(t, value[i], *pointer[i])
		}
	})
}
func Test_ByteValue(t *testing.T) {
	t.Run("byte", func(t *testing.T) {
		pointer := byte(42)

		value := ByteValue(&pointer)
		assert.Equal(t, pointer, value)
	})

	t.Run("byte/slice", func(t *testing.T) {
		p1 := byte(42)
		p2 := byte(69)
		p3 := byte(99)
		pointer := []*byte{&p1, &p2, &p3}

		value := ByteValueSlice(pointer)
		require.Len(t, value, len(pointer))
		for i := range pointer {
			assert.Equal(t, *pointer[i], value[i])
		}
	})

	t.Run("byte/map", func(t *testing.T) {
		p1 := byte(42)
		p2 := byte(69)
		p3 := byte(99)
		pointer := map[string]*byte{
			"foo": &p1, 
			"bar": &p2, 
			"baz": &p3,
		}

		value := ByteValueMap(pointer)
		require.Len(t, value, len(pointer))
		for i := range pointer {
			assert.Equal(t, *pointer[i], value[i])
		}
	})
}
func Test_Int(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		value := int(42)

		pointer := Int(value)
		assert.Equal(t, value, *pointer)
	})

	t.Run("int/slice", func(t *testing.T) {
		value := []int{int(42), int(69), int(99)}

		pointer := IntSlice(value)
		require.Len(t, pointer, len(value))
		for i := range value {
			assert.Equal(t, value[i], *pointer[i])
		}
	})

	t.Run("int/map", func(t *testing.T) {
		value := map[string]int{
			"foo": int(42), 
			"bar": int(69), 
			"baz": int(99),
		}

		pointer := IntMap(value)
		require.Len(t, pointer, len(value))
		for i := range value {
			assert.Equal(t, value[i], *pointer[i])
		}
	})
}
func Test_IntValue(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		pointer := int(42)

		value := IntValue(&pointer)
		assert.Equal(t, pointer, value)
	})

	t.Run("int/slice", func(t *testing.T) {
		p1 := int(42)
		p2 := int(69)
		p3 := int(99)
		pointer := []*int{&p1, &p2, &p3}

		value := IntValueSlice(pointer)
		require.Len(t, value, len(pointer))
		for i := range pointer {
			assert.Equal(t, *pointer[i], value[i])
		}
	})

	t.Run("int/map", func(t *testing.T) {
		p1 := int(42)
		p2 := int(69)
		p3 := int(99)
		pointer := map[string]*int{
			"foo": &p1, 
			"bar": &p2, 
			"baz": &p3,
		}

		value := IntValueMap(pointer)
		require.Len(t, value, len(pointer))
		for i := range pointer {
			assert.Equal(t, *pointer[i], value[i])
		}
	})
}
func Test_Int8(t *testing.T) {
	t.Run("int8", func(t *testing.T) {
		value := int8(42)

		pointer := Int8(value)
		assert.Equal(t, value, *pointer)
	})

	t.Run("int8/slice", func(t *testing.T) {
		value := []int8{int8(42), int8(69), int8(99)}

		pointer := Int8Slice(value)
		require.Len(t, pointer, len(value))
		for i := range value {
			assert.Equal(t, value[i], *pointer[i])
		}
	})

	t.Run("int8/map", func(t *testing.T) {
		value := map[string]int8{
			"foo": int8(42), 
			"bar": int8(69), 
			"baz": int8(99),
		}

		pointer := Int8Map(value)
		require.Len(t, pointer, len(value))
		for i := range value {
			assert.Equal(t, value[i], *pointer[i])
		}
	})
}
func Test_Int8Value(t *testing.T) {
	t.Run("int8", func(t *testing.T) {
		pointer := int8(42)

		value := Int8Value(&pointer)
		assert.Equal(t, pointer, value)
	})

	t.Run("int8/slice", func(t *testing.T) {
		p1 := int8(42)
		p2 := int8(69)
		p3 := int8(99)
		pointer := []*int8{&p1, &p2, &p3}

		value := Int8ValueSlice(pointer)
		require.Len(t, value, len(pointer))
		for i := range pointer {
			assert.Equal(t, *pointer[i], value[i])
		}
	})

	t.Run("int8/map", func(t *testing.T) {
		p1 := int8(42)
		p2 := int8(69)
		p3 := int8(99)
		pointer := map[string]*int8{
			"foo": &p1, 
			"bar": &p2, 
			"baz": &p3,
		}

		value := Int8ValueMap(pointer)
		require.Len(t, value, len(pointer))
		for i := range pointer {
			assert.Equal(t, *pointer[i], value[i])
		}
	})
}
func Test_Int16(t *testing.T) {
	t.Run("int16", func(t *testing.T) {
		value := int16(42)

		pointer := Int16(value)
		assert.Equal(t, value, *pointer)
	})

	t.Run("int16/slice", func(t *testing.T) {
		value := []int16{int16(42), int16(69), int16(99)}

		pointer := Int16Slice(value)
		require.Len(t, pointer, len(value))
		for i := range value {
			assert.Equal(t, value[i], *pointer[i])
		}
	})

	t.Run("int16/map", func(t *testing.T) {
		value := map[string]int16{
			"foo": int16(42), 
			"bar": int16(69), 
			"baz": int16(99),
		}

		pointer := Int16Map(value)
		require.Len(t, pointer, len(value))
		for i := range value {
			assert.Equal(t, value[i], *pointer[i])
		}
	})
}
func Test_Int16Value(t *testing.T) {
	t.Run("int16", func(t *testing.T) {
		pointer := int16(42)

		value := Int16Value(&pointer)
		assert.Equal(t, pointer, value)
	})

	t.Run("int16/slice", func(t *testing.T) {
		p1 := int16(42)
		p2 := int16(69)
		p3 := int16(99)
		pointer := []*int16{&p1, &p2, &p3}

		value := Int16ValueSlice(pointer)
		require.Len(t, value, len(pointer))
		for i := range pointer {
			assert.Equal(t, *pointer[i], value[i])
		}
	})

	t.Run("int16/map", func(t *testing.T) {
		p1 := int16(42)
		p2 := int16(69)
		p3 := int16(99)
		pointer := map[string]*int16{
			"foo": &p1, 
			"bar": &p2, 
			"baz": &p3,
		}

		value := Int16ValueMap(pointer)
		require.Len(t, value, len(pointer))
		for i := range pointer {
			assert.Equal(t, *pointer[i], value[i])
		}
	})
}
func Test_Int32(t *testing.T) {
	t.Run("int32", func(t *testing.T) {
		value := int32(42)

		pointer := Int32(value)
		assert.Equal(t, value, *pointer)
	})

	t.Run("int32/slice", func(t *testing.T) {
		value := []int32{int32(42), int32(69), int32(99)}

		pointer := Int32Slice(value)
		require.Len(t, pointer, len(value))
		for i := range value {
			assert.Equal(t, value[i], *pointer[i])
		}
	})

	t.Run("int32/map", func(t *testing.T) {
		value := map[string]int32{
			"foo": int32(42), 
			"bar": int32(69), 
			"baz": int32(99),
		}

		pointer := Int32Map(value)
		require.Len(t, pointer, len(value))
		for i := range value {
			assert.Equal(t, value[i], *pointer[i])
		}
	})
}
func Test_Int32Value(t *testing.T) {
	t.Run("int32", func(t *testing.T) {
		pointer := int32(42)

		value := Int32Value(&pointer)
		assert.Equal(t, pointer, value)
	})

	t.Run("int32/slice", func(t *testing.T) {
		p1 := int32(42)
		p2 := int32(69)
		p3 := int32(99)
		pointer := []*int32{&p1, &p2, &p3}

		value := Int32ValueSlice(pointer)
		require.Len(t, value, len(pointer))
		for i := range pointer {
			assert.Equal(t, *pointer[i], value[i])
		}
	})

	t.Run("int32/map", func(t *testing.T) {
		p1 := int32(42)
		p2 := int32(69)
		p3 := int32(99)
		pointer := map[string]*int32{
			"foo": &p1, 
			"bar": &p2, 
			"baz": &p3,
		}

		value := Int32ValueMap(pointer)
		require.Len(t, value, len(pointer))
		for i := range pointer {
			assert.Equal(t, *pointer[i], value[i])
		}
	})
}
func Test_Int64(t *testing.T) {
	t.Run("int64", func(t *testing.T) {
		value := int64(42)

		pointer := Int64(value)
		assert.Equal(t, value, *pointer)
	})

	t.Run("int64/slice", func(t *testing.T) {
		value := []int64{int64(42), int64(69), int64(99)}

		pointer := Int64Slice(value)
		require.Len(t, pointer, len(value))
		for i := range value {
			assert.Equal(t, value[i], *pointer[i])
		}
	})

	t.Run("int64/map", func(t *testing.T) {
		value := map[string]int64{
			"foo": int64(42), 
			"bar": int64(69), 
			"baz": int64(99),
		}

		pointer := Int64Map(value)
		require.Len(t, pointer, len(value))
		for i := range value {
			assert.Equal(t, value[i], *pointer[i])
		}
	})
}
func Test_Int64Value(t *testing.T) {
	t.Run("int64", func(t *testing.T) {
		pointer := int64(42)

		value := Int64Value(&pointer)
		assert.Equal(t, pointer, value)
	})

	t.Run("int64/slice", func(t *testing.T) {
		p1 := int64(42)
		p2 := int64(69)
		p3 := int64(99)
		pointer := []*int64{&p1, &p2, &p3}

		value := Int64ValueSlice(pointer)
		require.Len(t, value, len(pointer))
		for i := range pointer {
			assert.Equal(t, *pointer[i], value[i])
		}
	})

	t.Run("int64/map", func(t *testing.T) {
		p1 := int64(42)
		p2 := int64(69)
		p3 := int64(99)
		pointer := map[string]*int64{
			"foo": &p1, 
			"bar": &p2, 
			"baz": &p3,
		}

		value := Int64ValueMap(pointer)
		require.Len(t, value, len(pointer))
		for i := range pointer {
			assert.Equal(t, *pointer[i], value[i])
		}
	})
}
func Test_Uint8(t *testing.T) {
	t.Run("uint8", func(t *testing.T) {
		value := uint8(42)

		pointer := Uint8(value)
		assert.Equal(t, value, *pointer)
	})

	t.Run("uint8/slice", func(t *testing.T) {
		value := []uint8{uint8(42), uint8(69), uint8(99)}

		pointer := Uint8Slice(value)
		require.Len(t, pointer, len(value))
		for i := range value {
			assert.Equal(t, value[i], *pointer[i])
		}
	})

	t.Run("uint8/map", func(t *testing.T) {
		value := map[string]uint8{
			"foo": uint8(42), 
			"bar": uint8(69), 
			"baz": uint8(99),
		}

		pointer := Uint8Map(value)
		require.Len(t, pointer, len(value))
		for i := range value {
			assert.Equal(t, value[i], *pointer[i])
		}
	})
}
func Test_Uint8Value(t *testing.T) {
	t.Run("uint8", func(t *testing.T) {
		pointer := uint8(42)

		value := Uint8Value(&pointer)
		assert.Equal(t, pointer, value)
	})

	t.Run("uint8/slice", func(t *testing.T) {
		p1 := uint8(42)
		p2 := uint8(69)
		p3 := uint8(99)
		pointer := []*uint8{&p1, &p2, &p3}

		value := Uint8ValueSlice(pointer)
		require.Len(t, value, len(pointer))
		for i := range pointer {
			assert.Equal(t, *pointer[i], value[i])
		}
	})

	t.Run("uint8/map", func(t *testing.T) {
		p1 := uint8(42)
		p2 := uint8(69)
		p3 := uint8(99)
		pointer := map[string]*uint8{
			"foo": &p1, 
			"bar": &p2, 
			"baz": &p3,
		}

		value := Uint8ValueMap(pointer)
		require.Len(t, value, len(pointer))
		for i := range pointer {
			assert.Equal(t, *pointer[i], value[i])
		}
	})
}
func Test_Uint16(t *testing.T) {
	t.Run("uint16", func(t *testing.T) {
		value := uint16(42)

		pointer := Uint16(value)
		assert.Equal(t, value, *pointer)
	})

	t.Run("uint16/slice", func(t *testing.T) {
		value := []uint16{uint16(42), uint16(69), uint16(99)}

		pointer := Uint16Slice(value)
		require.Len(t, pointer, len(value))
		for i := range value {
			assert.Equal(t, value[i], *pointer[i])
		}
	})

	t.Run("uint16/map", func(t *testing.T) {
		value := map[string]uint16{
			"foo": uint16(42), 
			"bar": uint16(69), 
			"baz": uint16(99),
		}

		pointer := Uint16Map(value)
		require.Len(t, pointer, len(value))
		for i := range value {
			assert.Equal(t, value[i], *pointer[i])
		}
	})
}
func Test_Uint16Value(t *testing.T) {
	t.Run("uint16", func(t *testing.T) {
		pointer := uint16(42)

		value := Uint16Value(&pointer)
		assert.Equal(t, pointer, value)
	})

	t.Run("uint16/slice", func(t *testing.T) {
		p1 := uint16(42)
		p2 := uint16(69)
		p3 := uint16(99)
		pointer := []*uint16{&p1, &p2, &p3}

		value := Uint16ValueSlice(pointer)
		require.Len(t, value, len(pointer))
		for i := range pointer {
			assert.Equal(t, *pointer[i], value[i])
		}
	})

	t.Run("uint16/map", func(t *testing.T) {
		p1 := uint16(42)
		p2 := uint16(69)
		p3 := uint16(99)
		pointer := map[string]*uint16{
			"foo": &p1, 
			"bar": &p2, 
			"baz": &p3,
		}

		value := Uint16ValueMap(pointer)
		require.Len(t, value, len(pointer))
		for i := range pointer {
			assert.Equal(t, *pointer[i], value[i])
		}
	})
}
func Test_Uint32(t *testing.T) {
	t.Run("uint32", func(t *testing.T) {
		value := uint32(42)

		pointer := Uint32(value)
		assert.Equal(t, value, *pointer)
	})

	t.Run("uint32/slice", func(t *testing.T) {
		value := []uint32{uint32(42), uint32(69), uint32(99)}

		pointer := Uint32Slice(value)
		require.Len(t, pointer, len(value))
		for i := range value {
			assert.Equal(t, value[i], *pointer[i])
		}
	})

	t.Run("uint32/map", func(t *testing.T) {
		value := map[string]uint32{
			"foo": uint32(42), 
			"bar": uint32(69), 
			"baz": uint32(99),
		}

		pointer := Uint32Map(value)
		require.Len(t, pointer, len(value))
		for i := range value {
			assert.Equal(t, value[i], *pointer[i])
		}
	})
}
func Test_Uint32Value(t *testing.T) {
	t.Run("uint32", func(t *testing.T) {
		pointer := uint32(42)

		value := Uint32Value(&pointer)
		assert.Equal(t, pointer, value)
	})

	t.Run("uint32/slice", func(t *testing.T) {
		p1 := uint32(42)
		p2 := uint32(69)
		p3 := uint32(99)
		pointer := []*uint32{&p1, &p2, &p3}

		value := Uint32ValueSlice(pointer)
		require.Len(t, value, len(pointer))
		for i := range pointer {
			assert.Equal(t, *pointer[i], value[i])
		}
	})

	t.Run("uint32/map", func(t *testing.T) {
		p1 := uint32(42)
		p2 := uint32(69)
		p3 := uint32(99)
		pointer := map[string]*uint32{
			"foo": &p1, 
			"bar": &p2, 
			"baz": &p3,
		}

		value := Uint32ValueMap(pointer)
		require.Len(t, value, len(pointer))
		for i := range pointer {
			assert.Equal(t, *pointer[i], value[i])
		}
	})
}
func Test_Uint64(t *testing.T) {
	t.Run("uint64", func(t *testing.T) {
		value := uint64(42)

		pointer := Uint64(value)
		assert.Equal(t, value, *pointer)
	})

	t.Run("uint64/slice", func(t *testing.T) {
		value := []uint64{uint64(42), uint64(69), uint64(99)}

		pointer := Uint64Slice(value)
		require.Len(t, pointer, len(value))
		for i := range value {
			assert.Equal(t, value[i], *pointer[i])
		}
	})

	t.Run("uint64/map", func(t *testing.T) {
		value := map[string]uint64{
			"foo": uint64(42), 
			"bar": uint64(69), 
			"baz": uint64(99),
		}

		pointer := Uint64Map(value)
		require.Len(t, pointer, len(value))
		for i := range value {
			assert.Equal(t, value[i], *pointer[i])
		}
	})
}
func Test_Uint64Value(t *testing.T) {
	t.Run("uint64", func(t *testing.T) {
		pointer := uint64(42)

		value := Uint64Value(&pointer)
		assert.Equal(t, pointer, value)
	})

	t.Run("uint64/slice", func(t *testing.T) {
		p1 := uint64(42)
		p2 := uint64(69)
		p3 := uint64(99)
		pointer := []*uint64{&p1, &p2, &p3}

		value := Uint64ValueSlice(pointer)
		require.Len(t, value, len(pointer))
		for i := range pointer {
			assert.Equal(t, *pointer[i], value[i])
		}
	})

	t.Run("uint64/map", func(t *testing.T) {
		p1 := uint64(42)
		p2 := uint64(69)
		p3 := uint64(99)
		pointer := map[string]*uint64{
			"foo": &p1, 
			"bar": &p2, 
			"baz": &p3,
		}

		value := Uint64ValueMap(pointer)
		require.Len(t, value, len(pointer))
		for i := range pointer {
			assert.Equal(t, *pointer[i], value[i])
		}
	})
}
func Test_Float32(t *testing.T) {
	t.Run("float32", func(t *testing.T) {
		value := float32(42)

		pointer := Float32(value)
		assert.Equal(t, value, *pointer)
	})

	t.Run("float32/slice", func(t *testing.T) {
		value := []float32{float32(42), float32(69), float32(99)}

		pointer := Float32Slice(value)
		require.Len(t, pointer, len(value))
		for i := range value {
			assert.Equal(t, value[i], *pointer[i])
		}
	})

	t.Run("float32/map", func(t *testing.T) {
		value := map[string]float32{
			"foo": float32(42), 
			"bar": float32(69), 
			"baz": float32(99),
		}

		pointer := Float32Map(value)
		require.Len(t, pointer, len(value))
		for i := range value {
			assert.Equal(t, value[i], *pointer[i])
		}
	})
}
func Test_Float32Value(t *testing.T) {
	t.Run("float32", func(t *testing.T) {
		pointer := float32(42)

		value := Float32Value(&pointer)
		assert.Equal(t, pointer, value)
	})

	t.Run("float32/slice", func(t *testing.T) {
		p1 := float32(42)
		p2 := float32(69)
		p3 := float32(99)
		pointer := []*float32{&p1, &p2, &p3}

		value := Float32ValueSlice(pointer)
		require.Len(t, value, len(pointer))
		for i := range pointer {
			assert.Equal(t, *pointer[i], value[i])
		}
	})

	t.Run("float32/map", func(t *testing.T) {
		p1 := float32(42)
		p2 := float32(69)
		p3 := float32(99)
		pointer := map[string]*float32{
			"foo": &p1, 
			"bar": &p2, 
			"baz": &p3,
		}

		value := Float32ValueMap(pointer)
		require.Len(t, value, len(pointer))
		for i := range pointer {
			assert.Equal(t, *pointer[i], value[i])
		}
	})
}
func Test_Float64(t *testing.T) {
	t.Run("float64", func(t *testing.T) {
		value := float64(42)

		pointer := Float64(value)
		assert.Equal(t, value, *pointer)
	})

	t.Run("float64/slice", func(t *testing.T) {
		value := []float64{float64(42), float64(69), float64(99)}

		pointer := Float64Slice(value)
		require.Len(t, pointer, len(value))
		for i := range value {
			assert.Equal(t, value[i], *pointer[i])
		}
	})

	t.Run("float64/map", func(t *testing.T) {
		value := map[string]float64{
			"foo": float64(42), 
			"bar": float64(69), 
			"baz": float64(99),
		}

		pointer := Float64Map(value)
		require.Len(t, pointer, len(value))
		for i := range value {
			assert.Equal(t, value[i], *pointer[i])
		}
	})
}
func Test_Float64Value(t *testing.T) {
	t.Run("float64", func(t *testing.T) {
		pointer := float64(42)

		value := Float64Value(&pointer)
		assert.Equal(t, pointer, value)
	})

	t.Run("float64/slice", func(t *testing.T) {
		p1 := float64(42)
		p2 := float64(69)
		p3 := float64(99)
		pointer := []*float64{&p1, &p2, &p3}

		value := Float64ValueSlice(pointer)
		require.Len(t, value, len(pointer))
		for i := range pointer {
			assert.Equal(t, *pointer[i], value[i])
		}
	})

	t.Run("float64/map", func(t *testing.T) {
		p1 := float64(42)
		p2 := float64(69)
		p3 := float64(99)
		pointer := map[string]*float64{
			"foo": &p1, 
			"bar": &p2, 
			"baz": &p3,
		}

		value := Float64ValueMap(pointer)
		require.Len(t, value, len(pointer))
		for i := range pointer {
			assert.Equal(t, *pointer[i], value[i])
		}
	})
}
