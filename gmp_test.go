package main

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestProblem1(t *testing.T) {
	runtime.GOMAXPROCS(1)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	// 10 10 10 10 10
	time.Sleep(2 * time.Second)
}

func TestProblem2(t *testing.T) {
	runtime.GOMAXPROCS(1)
	for i := 0; i < 10; i++ {
		i := i
		go func() {
			fmt.Println(i)
		}()
	}
	// 9 0 1 2 3
	time.Sleep(2 * time.Second)
}
func TestProblem3(t *testing.T) {
	var x int
	threads := runtime.GOMAXPROCS(0)
	fmt.Println("threads = ", threads)
	for i := 0; i < 8; i++ {
		go func() {
			for {
				x = x + 1
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Println("x =", x)
}

func TestProblem4(t *testing.T) {
	var x, y int
	//G1
	go func() {
		x = 1    // A1
		t.Log(y) // A2
	}()
	//G2
	go func() {
		y = 1    // A3
		t.Log(x) // A4
	}()
	time.Sleep(time.Second)
}
