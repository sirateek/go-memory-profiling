package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http/pprof"
)

func main() {
	server := gin.Default()

	ctx := context.Background()
	newCtx, cancel := context.WithCancel(ctx)

	handler := Handler{
		ctx: newCtx,
	}
	server.POST("/example", handler.HandleRequest)

	server.GET("/debug/pprof", gin.WrapF(pprof.Index))
	server.GET("/debug/heap", gin.WrapH(pprof.Handler("heap")))

	server.Run(":8080")

	// Gracefully Shutting Down.
	cancel()

}
