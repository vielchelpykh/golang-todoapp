package domains

// для верного декодирования входящих данных для изменений сущностей
type Nullable[T any] struct {
	Value *T
	Set   bool
}
