package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
)

type ratelimitByCountBuilder struct {
	count int64
}

func newRateLimit(count int64) *ratelimitByCountBuilder {
	//todo 启动协程清空count
	builder := &ratelimitByCountBuilder{
		count: count,
	}

	go func() {
		ticker := time.NewTicker(time.Second)
		for {
			select {
			case <-ticker.C:
				atomic.StoreInt64(&builder.count, count)
			}
		}
	}()
	return builder
}

func (builder *ratelimitByCountBuilder) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if atomic.LoadInt64(&builder.count) <= 0 {
			fmt.Printf("rejection req!\n")
			ctx.Abort()
			return
		}

		atomic.AddInt64(&builder.count, -1)
	}
}

func main() {
	server := gin.Default()
	rate := newRateLimit(3)
	server.GET("/", rate.Build(), func(ctx *gin.Context) { ctx.String(http.StatusOK, "hi") })
	server.Run()
}
