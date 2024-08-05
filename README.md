# ptr

Golang value <-> pointer utils.  
Inspired by [aws pointer utils](https://github.com/aws/aws-sdk-go/blob/main/aws/convert_types.go).  
The goal is to avoid using AWS pointer utilities for convenience, spreading AWS dependencies where they are not needed.

## Example

### Value -> Pointer
```go
// before
someString := "foo"
someSDK.SomeMethod(&someString)

// after
someSDK.SomeMethod(ptr.String("foo"))
```

### Pointer -> Value
```go
// before
strPtr := someSDK.GetStringPtr()
var str string
if strPtr != nil {
  str = *strPtr
} 

// after
strPtr := someSDK.GetStringPtr()
str := ptr.StrinValue(strPtr)
```

## Generic core funcs
Core functions use generics `Ptr[T any](v T) *t` and `Value[T any](p *T) T`.
So you can easily use it with structs, slices and maps.

It also offers slice and maps transformations like the aws utility.
```go
func PtrSlice[T any](v []T) []*T
func PtrMap[K comparable, T any](v map[K]T) map[K]*T
func ValueSlice[T any](p []*T) []T
func ValueMap[K comparable, T any](v map[K]*T) map[K]T
```

It offers out of the box type wrappers for:
- string
- byte
- bool
- int, int8, int16, int32 and int64
- uint, uint8, uint16, uint32 and uint64
- float32 and float64
- time.Time

#### Development guide
For sake of avoiding writing the same lines of code for both `ptr.go` and `ptr_test.go` to create the type wrappers, I created the `./generator.go` to generate the type boilerplate for both source and test.

To generaten run:
```shell
cd geneator
go run main.go
```