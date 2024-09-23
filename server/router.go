package server

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Start(ctx context.Context, h *handler, port uint64) {
	r := gin.Default()

	r.GET("/ping", h.pingHandler)
	r.POST("/accounts", h.createAccountHandler)
	r.GET("/accounts/:accountId", h.getAccountHandler)
	r.POST("/transactions", h.createTransactionHandler)
	r.Run(fmt.Sprintf(":%d", port))
}
