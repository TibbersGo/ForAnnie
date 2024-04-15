package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//for i := 0; i <= 10; i++ {
	//	go func(gid int) {
	//		n := 0
	//		for {
	//			fmt.Println(time.Now().Format("2006-01-02 15:04:05"), gid, n)
	//			time.Sleep(time.Second)
	//		}
	//	}(i)
	//}
	//go func() {
	//	arr := 0
	//	p := uintptr(unsafe.Pointer(&arr))
	//	myfunc1(p)
	//}()
	//a := [3]int{1, 2, 3}
	//a[0] = 4
	//fmt.Println(a)

	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Println("loop.")
		}()
		fmt.Println(i)
	}

}

func myfunc1(p uintptr) {
	arr := (*int)(unsafe.Pointer(p))
	*arr = 1
	fmt.Println(*arr)
	go myfunc2()
	fmt.Println(*arr)
}

func myfunc2() {
	fmt.Println("myfunc2")
	myfunc3()
}

func myfunc3() {
	var p uintptr = 0
	arr := (*int)(unsafe.Pointer(p))
	*arr = 1
	fmt.Println(*arr)
}
