package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type limitBySlidingWindowBuilder struct {
	windowSize  time.Duration //窗口大小
	maxRequests int           //最大请求数量
	requests    []time.Time
	lock        sync.Mutex
}

func newBuilder(size time.Duration, max int) *limitBySlidingWindowBuilder {
	return &limitBySlidingWindowBuilder{
		windowSize:  size,
		maxRequests: max,
		requests:    make([]time.Time, 0, max),
	}
}

func (limit *limitBySlidingWindowBuilder) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !limit.allowRequest() {
			ctx.Abort()
			return
		}
	}
}

func (limit *limitBySlidingWindowBuilder) allowRequest() bool {
	currentTime := time.Now()
	limit.lock.Lock()
	defer limit.lock.Unlock()

	//清理过期请求
	if len(limit.requests) > 0 && currentTime.Sub(limit.requests[0]) > limit.windowSize {
		limit.requests = limit.requests[1:]
	}

	//检查请求数
	if len(limit.requests) >= limit.maxRequests {
		fmt.Println("reject")
		return false
	}

	limit.requests = append(limit.requests, currentTime)

	return true
}

// todo 自定义动态调整窗口大小
func (limit *limitBySlidingWindowBuilder) update(t time.Duration) {
	limit.lock.Lock()
	defer limit.lock.Unlock()
	limit.windowSize = t

}

func main() {
	server := gin.Default()
	server.GET("/", newBuilder(time.Second, 1).Build(), func(ctx *gin.Context) { ctx.String(200, "hi") })
	server.Run()
}
