package zmodels

type ZRequest[T any] struct {
	Url         string
	Headers     map[string]string
	Body        T
	UrlParams   map[string]string
	QueryParams map[string][]string
}
