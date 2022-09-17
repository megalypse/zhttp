package zmodels

import "net/http"

type ZResponse[T any] struct {
	Content      *T
	Response     *http.Response
	IsSuccess    bool
	ErrorMessage string
}

func MakeFailResponse[T any](message string, httpResponse *http.Response) ZResponse[T] {
	return makeResponse[T](nil, httpResponse, false, message)
}

func makeResponse[T any](
	content *T,
	response *http.Response,
	isSuccess bool,
	errorMessage string,
) ZResponse[T] {
	return ZResponse[T]{
		Content:      content,
		Response:     response,
		IsSuccess:    isSuccess,
		ErrorMessage: errorMessage,
	}
}
