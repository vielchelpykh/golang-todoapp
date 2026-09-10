package domain

// для верного декодирования входящих данных при изменени сущностей
/*
Nullable needs to specify:
  - field not provided
  - field provided: value
  - field provided: null
*/
type Nullable[T any] struct {
	Value *T
	Set   bool
}
