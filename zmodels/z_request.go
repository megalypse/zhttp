package zmodels

type ZRequest[T any] struct {
	Url         string
	Headers     []ZHeader
	Body        T
	UrlParams   []Param
	QueryParams []Param
}
