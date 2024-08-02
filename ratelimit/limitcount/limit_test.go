package main

import (
	"log"
	"net/http"
	"sync"
	"testing"
)

func TestRatelimit(t *testing.T) {
	var g sync.WaitGroup
	g.Add(4)
	for i := 0; i < 4; i++ {
		go func() {
			_, err := http.Get("http://localhost:8080")
			if err != nil {
				log.Fatal(err)
			}
			g.Done()
		}()
	}

	g.Wait()
}
