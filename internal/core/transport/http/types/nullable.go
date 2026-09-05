package core_http_types

import (
	"encoding/json"

	"github.com/vielchelpykh/golang-todoapp/internal/core/domains"
)

type Nullable[T any] struct {
	domains.Nullable[T]
}

// если вызван, то поле было передано, то есть Set = true
func (n *Nullable[T]) UnmarshalJSON(b []byte) error {
	n.Set = true

	if string(b) == "null" {
		n.Value = nil

		return nil
	}

	var value T
	if err := json.Unmarshal(b, &value); err != nil {
		return err
	}

	n.Value = &value
	return nil
}

func (n *Nullable[T]) ToDomain() domains.Nullable[T] {
	return domains.Nullable[T]{
		Value: n.Value,
		Set:   n.Set,
	}
}
