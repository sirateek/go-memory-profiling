package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

const (
	HeaderReqIDKey = "X-REQ-ID"
)

type Handler struct {
	ctx context.Context
}

func (h *Handler) HandleRequest(g *gin.Context) {
	reqID := g.GetHeader(HeaderReqIDKey)
	ctx := context.WithValue(h.ctx, HeaderReqIDKey, reqID)

	var wg sync.WaitGroup
	wg.Add(2)
	go func(innerContext context.Context) {
		// Do Something
		fmt.Println("First Routine: ", innerContext.Value(HeaderReqIDKey))
		wg.Done()
	}(ctx)

	go func(innerContext context.Context) {
		// Do Something
		fmt.Println("Second Routine: ", innerContext.Value(HeaderReqIDKey))
		wg.Done()
	}(ctx)

	wg.Wait()
	g.String(http.StatusOK, reqID)
}
