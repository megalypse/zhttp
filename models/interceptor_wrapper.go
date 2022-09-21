package models

type InterceptorsWrapper[Request any, Response any] struct {
	RequestInterceptors  []Interceptor[Request, ZRequest[Request]]
	ResponseInterceptors []Interceptor[Response, ZResponse[Response]]
}
