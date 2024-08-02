package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type leakybucket struct {
	rate       float64 //漏桶速度 请求数/秒
	size       int
	water      int   //当前水量
	lastLeakMs int64 //上次漏水时间戳
	lock       sync.Mutex
}

func newleak(rate float64, size int) *leakybucket {
	return &leakybucket{
		rate:       rate,
		size:       size,
		water:      0,
		lastLeakMs: time.Now().Unix(),
	}
}

func (limit *leakybucket) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !limit.allow() {
			ctx.Abort()
			return
		}
	}
}

func (limit *leakybucket) allow() bool {
	now := time.Now().Unix()

	limit.lock.Lock()
	defer limit.lock.Unlock()

	//之前漏出的水量
	leakAmount := int(float64(now-limit.lastLeakMs) / 1000 * limit.rate)

	if leakAmount > 0 {
		if leakAmount > limit.water {
			limit.water = 0
		} else {
			limit.water -= leakAmount
		}
	}

	//计算当前是否超过容量
	if limit.water > limit.size {
		limit.water--
		fmt.Println("reject")
		return false
	}

	limit.water++

	limit.lastLeakMs = now
	return true
}

func main() {
	s := gin.Default()
	s.GET("/", newleak(1, 4).Build(), func(ctx *gin.Context) {
		fmt.Print("hi")
	})
	s.Run()

}
