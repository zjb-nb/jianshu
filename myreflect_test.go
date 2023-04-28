package main

import (
	"reflect"
	"testing"
)

func TestChange(t *testing.T) {
	var x int64 = 1
	v := reflect.ValueOf(&x)
	// v := reflect.ValueOf(x)
	v.Elem().SetInt(2)
	t.Log(x)
}
