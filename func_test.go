package main

import (
	"fmt"
	"testing"
)

func TestParam(t *testing.T) {
	a := 30
	b := &a
	fmt.Printf("in main a address:%d=%p,b address:%d,%p\n", a, &a, *b, &b)
	myfunction(a, b)
}

func myfunction(i int, j *int) {
	fmt.Printf("in myfunction i address:%d=%p,j address:%d,%p\n", i, &i, *j, &j)
}

func TestDefer(t *testing.T) {
	var a, b int
	b = incr(a)
	t.Log(b)
}

func incr(a int) int {
	var b int
	defer func() {
		a++
		b++
	}()
	a++
	b = a
	return b
}

func incr2(a int) (b int) {

	defer func() {
		a++
		b++
	}()
	a++
	b = a
	return b
}
