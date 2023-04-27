package main

import (
	"fmt"
	"testing"
)

type Languge interface {
	sayHello()
	// code()
}
type Go struct{}

func (g Go) sayHello() {
	fmt.Println("hello world go")
}

//	func (g *Go) code() {
//		fmt.Println("code go")
//	}
var _ Languge = (*Go)(nil)

func sayHello(l Languge) {
	l.sayHello()
	// l.code()
}
func TestInterface(t *testing.T) {
	g := new(Go)
	sayHello(g)
}

func TestDiff(t *testing.T) {
	var g Languge = Go{}
	// g.code()
	g.sayHello()
}

func TestCompareNil(t *testing.T) {
	var g Languge
	t.Log(g == nil)
	fmt.Printf("g:%T,%v\n", g, g)

	var g1 *Go
	t.Log(g1 == nil)

	g = g1
	t.Log(g == nil)
	fmt.Printf("g:%T,%v", g, g)
}

func TestCompareNil2(t *testing.T) {
	var p func() Languge
	p = func() Languge {
		var g *Go = nil
		return g
	}
	g1 := p()
	t.Log(g1)
	t.Log(g1 == nil)
}

func TestCompareNil3(t *testing.T) {
	var p func(v interface{}) bool
	p = func(v interface{}) bool {
		return v == nil
	}
	var g *Go
	var te interface{}
	t.Log(te == nil)
	t.Log(g == nil)
	t.Log(p(g))
}
