package main

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestTypeAlise(t *testing.T) {
	type myint int
	var a int = 1
	var b myint = 2
	// t.Log(a == b)
	t.Log(a == (int)(b))
}

func TestStrLen(t *testing.T) {
	var s string = "zjb"
	t.Log(s[1])
	// s[1] = 'g'
	s = "www"
}

func TestStrLen1(t *testing.T) {
	var a string = "yes"
	var b string = "我爱你"

	fmt.Println(len(a)) //3
	fmt.Println(len(b)) //9

	fmt.Println(unsafe.Sizeof(a))   //16
	fmt.Println(unsafe.Sizeof(b))   //16
	fmt.Println(unsafe.Sizeof("我")) //16

	var s string = "中国"
	t.Log(len(s))
	for _, c := range s {
		fmt.Printf("%x\n", c)
	}
}

func TestRuneAndByte(t *testing.T) {
	var s string = "go中国"
	t.Log(len(s))
	t.Log([]rune(s))
	t.Log([]byte(s))

	t.Log(unsafe.Sizeof('y'))
	t.Log(unsafe.Sizeof((byte)('y')))
}

func TestCompareStr1(t *testing.T) {
	s1 := "12345"
	// s2 := "2"
	s2 := "12"
	t.Log(s1 > s2)
	t.Log([]byte(s1), []byte(s2))
}

func TestCompareStr2(t *testing.T) {
	s1 := "零"
	s2 := "一"
	s3 := "二"

	t.Log(s1 > s2, []byte(s1), []byte(s2))
	t.Log(s3 > s2, []byte(s3), []byte(s2))
	t.Log(s3 > s1, []byte(s3), []byte(s1))
}
