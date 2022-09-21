package models

type Interceptor[V any, T ZRequest[V] | ZResponse[V]] func(payload T) (bool, T)
