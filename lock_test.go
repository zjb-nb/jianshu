package main

import (
	"sync"
	"testing"
	"time"
)

type MyData struct {
	id int
	mu sync.Mutex
}

func TestLockAgain(t *testing.T) {
	data := MyData{id: 1}
	data.mu.Lock()
	data.mu.Lock()
	data.mu.Unlock()
	data.mu.Unlock()
}

func TestUnLockAgain(t *testing.T) {
	data := MyData{id: 0}
	data.mu.Lock()
	defer data.mu.Unlock()
	data.mu.Unlock()
}

func TestCopyMutex(t *testing.T) {
	data := MyData{id: 0}
	go func() {
		data.mu.Lock()
		defer data.mu.Unlock()
		time.Sleep(time.Hour)
	}()
	time.Sleep(time.Second)
	go func(mydata MyData) {
		mydata.mu.Unlock()
		mydata.id = 2
		t.Log(mydata.id)
		time.Sleep(time.Hour)
	}(data)
	time.Sleep(time.Second * 5)
}

func TestDeadlock(t *testing.T) {
	d1 := MyData{id: 1}
	d2 := MyData{id: 2}
	go func() {
		d1.mu.Lock()
		defer d1.mu.Unlock()
		time.Sleep(time.Second)
		t.Log("g1-d1.id:", d1.id)
		d2.mu.Lock()
		t.Log("g1-d2.id:", d2.id)
		d2.mu.Unlock()
	}()
	go func() {
		d2.mu.Lock()
		defer d2.mu.Unlock()
		time.Sleep(time.Second)
		t.Log("g2-d2.id:", d2.id)
		d1.mu.Lock()
		t.Log("g2-d1.id:", d1.id)
		d1.mu.Unlock()
	}()
	var ch = make(chan int)
	<-ch
}

func TestDoubleCheck(t *testing.T) {
	type safeData struct {
		values map[int]int
		lock   sync.RWMutex
	}

	data := safeData{values: make(map[int]int)}
	chang := func(key int, val int) int {
		data.lock.RLock()
		old, ok := data.values[key]
		data.lock.RUnlock()
		if ok {
			return old
		}

		data.lock.Lock()
		defer data.lock.Unlock()
		data.values[key] = val
		return val
	}
	chang2 := func(key int, val int) int {
		data.lock.RLock()
		old, ok := data.values[key]
		data.lock.RUnlock()
		if ok {
			return old
		}

		data.lock.Lock()
		defer data.lock.Unlock()
		//这里添加double-check
		old, ok = data.values[key]
		if ok {
			return old
		}
		data.values[key] = val
		return val
	}
	go chang(0, 1)
	go chang2(0, 2)
}
