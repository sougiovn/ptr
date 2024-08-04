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
		myVal := teststruct{foo: "baz", bar: 42}

		myPtr := Ptr(myVal)
		assert.Equal(t, myVal, *myPtr)
	})

	t.Run("slice/string", func(t *testing.T) {
		myVal := []string{"foo", "bar"}

		myPtr := Ptr(myVal)
		assert.Equal(t, myVal, *myPtr)
	})

	t.Run("slice/struct", func(t *testing.T) {
		myVal := []teststruct{
			{"foo", 42},
			{"bar", 69},
		}

		myPtr := Ptr(myVal)
		assert.Equal(t, myVal, *myPtr)
	})
}

func Test_PtrSlice(t *testing.T) {
	t.Run("string[]", func(t *testing.T) {
		myVal := []string{"foo", "bar", "baz"}

		myPtr := PtrSlice(myVal)
		require.Len(t, myPtr, len(myVal))
		for i := range myVal {
			assert.Equal(t, myVal[i], *myPtr[i])
		}
	})

	t.Run("struct[]", func(t *testing.T) {
		myVal := []teststruct{
			{"foo", 42},
			{"bar", 69},
			{"baz", 99},
		}

		myPtr := PtrSlice(myVal)
		require.Len(t, myPtr, len(myVal))
		for i := range myVal {
			assert.Equal(t, myVal[i], *myPtr[i])
		}
	})
}

func Test_PtrMap(t *testing.T) {
	t.Run("[string]string", func(t *testing.T) {
		myVal := map[string]string{
			"foo": "foo",
			"bar": "bar",
			"baz": "baz",
		}

		myPtr := PtrMap(myVal)
		require.Len(t, myPtr, len(myVal))
		for k := range myVal {
			assert.Equal(t, myVal[k], *myPtr[k])
		}
	})

	t.Run("[int]string", func(t *testing.T) {
		myVal := map[int]string{
			42: "foo",
			69: "bar",
			99: "baz",
		}

		myPtr := PtrMap(myVal)
		require.Len(t, myPtr, len(myVal))
		for k := range myVal {
			assert.Equal(t, myVal[k], *myPtr[k])
		}
	})

	t.Run("[bool]int", func(t *testing.T) {
		myVal := map[bool]int{
			true:  42,
			false: 69,
		}

		myPtr := PtrMap(myVal)
		require.Len(t, myPtr, len(myVal))
		for k := range myVal {
			assert.Equal(t, myVal[k], *myPtr[k])
		}
	})

	t.Run("[string]struct", func(t *testing.T) {
		myVal := map[string]teststruct{
			"foo": {"foo", 42},
			"bar": {"bar", 69},
			"baz": {"baz", 420},
		}

		myPtr := PtrMap(myVal)
		require.Len(t, myPtr, len(myVal))
		for k := range myVal {
			assert.Equal(t, myVal[k], *myPtr[k])
		}
	})
}

func Test_Value(t *testing.T) {
	t.Run("struct", func(t *testing.T) {
		myPtr := &teststruct{foo: "baz", bar: 42}

		myValue := Value(myPtr)
		assert.Equal(t, *myPtr, myValue)
		assert.Equal(t, myPtr, &myValue)
	})

	t.Run("slice/string", func(t *testing.T) {
		myPtr := &[]string{"foo", "bar"}

		myValue := Value(myPtr)
		assert.Equal(t, *myPtr, myValue)
		assert.Equal(t, myPtr, &myValue)
	})

	t.Run("slice/struct", func(t *testing.T) {
		myPtr := &[]teststruct{
			{"foo", 42},
			{"bar", 69},
		}

		myValue := Value(myPtr)
		assert.Equal(t, *myPtr, myValue)
		assert.Equal(t, myPtr, &myValue)
	})

	t.Run("nil", func(t *testing.T) {
		var myPtr *teststruct

		myValue := Value(myPtr)
		assert.Equal(t, teststruct{}, myValue)
	})
}

func Test_ValueSlice(t *testing.T) {
	t.Run("string[]", func(t *testing.T) {
		foo := "foo"
		bar := "bar"
		baz := "baz"
		myPtr := []*string{&foo, &bar, &baz}

		myVal := ValueSlice(myPtr)
		require.Len(t, myVal, len(myPtr))
		for i := range myPtr {
			assert.Equal(t, *myPtr[i], myVal[i])
		}
	})

	t.Run("struct[]", func(t *testing.T) {
		foo := teststruct{"foo", 42}
		bar := teststruct{"bar", 69}
		baz := teststruct{"baz", 99}
		myPtr := []*teststruct{&foo, &bar, &baz}

		myVal := ValueSlice(myPtr)
		require.Len(t, myVal, len(myPtr))
		for i := range myPtr {
			assert.Equal(t, *myPtr[i], myVal[i])
		}
	})
}

func Test_ValueMap(t *testing.T) {
	t.Run("[string]string", func(t *testing.T) {
		foo := "foo"
		bar := "bar"
		baz := "baz"
		myPtr := map[string]*string{
			"foo": &foo,
			"bar": &bar,
			"baz": &baz,
		}

		myVal := ValueMap(myPtr)
		require.Len(t, myVal, len(myVal))
		for k := range myPtr {
			assert.Equal(t, *myPtr[k], myVal[k])
		}
	})

	t.Run("[int]string", func(t *testing.T) {
		foo := "foo"
		bar := "bar"
		baz := "baz"
		myPtr := map[int]*string{
			42: &foo,
			69: &bar,
			99: &baz,
		}

		myVal := ValueMap(myPtr)
		require.Len(t, myVal, len(myVal))
		for k := range myPtr {
			assert.Equal(t, *myPtr[k], myVal[k])
		}
	})

	t.Run("[bool]int", func(t *testing.T) {
		b1 := 42
		b2 := 69
		myPtr := map[bool]*int{
			true:  &b1,
			false: &b2,
		}

		myVal := ValueMap(myPtr)
		require.Len(t, myVal, len(myVal))
		for k := range myPtr {
			assert.Equal(t, *myPtr[k], myVal[k])
		}
	})

	t.Run("[string]struct", func(t *testing.T) {
		foo := teststruct{"foo", 42}
		bar := teststruct{"bar", 69}
		baz := teststruct{"baz", 99}
		myPtr := map[string]*teststruct{
			"foo": &foo,
			"bar": &bar,
			"baz": &baz,
		}

		myVal := ValueMap(myPtr)
		require.Len(t, myVal, len(myVal))
		for k := range myPtr {
			assert.Equal(t, *myPtr[k], myVal[k])
		}
	})
}
