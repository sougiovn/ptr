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
