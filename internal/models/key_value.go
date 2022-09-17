package models

type KeyValue[T any] struct {
	Key   string
	Value T
}
