package main

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestNil(t *testing.T) {
	var a int = 0
	t.Log(&a == nil)
	type student struct{}
	// var s student
	// t.Log(s==nil)
}

func TestNilandAnotherType(t *testing.T) {
	type A interface{}
	type B struct{}
	var a *B //a是一个指针
	var m map[int]int
	var ch chan int
	var f func()
	var s []int

	t.Log(a == nil)
	t.Log(m == nil)
	t.Log(ch == nil)
	t.Log(f == nil)
	t.Log(s == nil)

}

func TestNiltoAnotherType(t *testing.T) {
	type B struct{ name string }
	var a *B //a是一个指针

	t.Log(a == nil)       //零值指针 true
	t.Log(a == (*B)(nil)) //true
}

func TestChangeNil(t *testing.T) {
	var nil int = 1
	a := 1
	t.Log(a == nil)
}

func TestStateNil(t *testing.T) {
	// a := nil
	// t.Log(a)
}

func TestSizeofNil(t *testing.T) {
	var p *struct{} = nil
	fmt.Println(unsafe.Sizeof(p)) // 8

	var s []int = nil
	fmt.Println(unsafe.Sizeof(s)) // 24

	var m map[int]bool = nil
	fmt.Println(unsafe.Sizeof(m)) // 8

	var c chan string = nil
	fmt.Println(unsafe.Sizeof(c)) // 8

	var f func() = nil
	fmt.Println(unsafe.Sizeof(f)) // 8

	var i interface{} = nil
	fmt.Println(unsafe.Sizeof(i)) // 16
}
