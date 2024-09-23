package errors

import (
	"errors"
	"net/http"
)

var (
	ErrDocumentNumberEmpty  = errors.New("document number is empty")
	ErrIDEmpty              = errors.New("request id is empty")
	ErrNoRows               = errors.New("no rows found")
	ErrInvalidAccountID     = errors.New("invalid account id")
	ErrInvalidOperationType = errors.New("invalid operation type")
)

func GetHttpError(err error) int {
	switch err {
	case ErrDocumentNumberEmpty, ErrIDEmpty:
		return http.StatusBadRequest
	case ErrNoRows:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
