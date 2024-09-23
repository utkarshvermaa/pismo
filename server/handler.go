package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/projects/sys-des/txn-routine/internal/accounts"
	accPayload "github.com/projects/sys-des/txn-routine/internal/accounts/payload"
	"github.com/projects/sys-des/txn-routine/internal/transactions"
	txnPayload "github.com/projects/sys-des/txn-routine/internal/transactions/payload"
	"github.com/projects/sys-des/txn-routine/pkg/errors"
)

type handler struct {
	acc accounts.Interface
	txn transactions.Interface
}

func Handler(
	acc accounts.Interface,
	txn transactions.Interface,
) *handler {
	return &handler{
		acc: acc,
		txn: txn,
	}
}

func (h *handler) pingHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (h *handler) createAccountHandler(ctx *gin.Context) {
	var r accPayload.CreateRequest
	if err := ctx.ShouldBindJSON(&r); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := h.acc.CreateAccount(ctx, &r)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":      resp.ID,
		"message": "account created",
	})
}

func (h *handler) getAccountHandler(ctx *gin.Context) {
	parsedAccId := ctx.Param("accountId")
	accountID, err := strconv.ParseUint(parsedAccId, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	r := accPayload.GetRequest{
		ID: accountID,
	}

	resp, err := h.acc.GetAccount(ctx, &r)
	if err != nil {
		ctx.JSON(errors.GetHttpError(err), gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *handler) createTransactionHandler(ctx *gin.Context) {
	var r txnPayload.CreateRequest

	if err := ctx.ShouldBindJSON(&r); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := h.txn.CreateTransaction(ctx, &r)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":      resp.ID,
		"message": "transaction created",
	})
}
