package unsafe_cmd

import (
	"fmt"
	"reflect"
	"unsafe"
)

type MyT struct {
	i32 int32
	i64 int64
}

func (myT *MyT) GetI32() {
	fmt.Printf("value: %v", myT.i32)
}

func (myT *MyT) GetI64() {
	fmt.Printf("value: %v", myT.i64)
}

func UnsafeMain() {
	my := &MyT{}
	i32 := (*int32)(unsafe.Pointer(my))
	*i32 = 187

	i64 := (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(my)) + uintptr(4)))
	i644 := (*int64)(unsafe.Add(unsafe.Pointer(my), 8))

	fmt.Println(unsafe.Offsetof(my.i32))
	fmt.Println(unsafe.Offsetof(my.i64))
	fmt.Println(reflect.ValueOf(i644))
	fmt.Println(reflect.ValueOf(&my.i32))
	fmt.Println(reflect.ValueOf(&my.i64))
	*i64 = 100001
	my.GetI32()
	fmt.Println()
	my.GetI64()
	fmt.Println()

	*i644 = 100001
	my.GetI64()
	fmt.Println()
}
