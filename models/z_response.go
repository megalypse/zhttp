package models

import "net/http"

type ZResponse[T any] struct {
	Content      *T
	Response     *http.Response
	IsSuccess    bool
	ErrorMessage string
}
