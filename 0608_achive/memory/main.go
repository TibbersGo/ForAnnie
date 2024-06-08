package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

type Student struct {
	a int8
	d int32
	b int8
	c int8
}
type A struct {
	a int8
	b string
	c int8
	d int8
	e int8
}

func main() {
	var t Student
	fmt.Println(unsafe.Sizeof(t)) //输出：5

	runtime.GC()
}
