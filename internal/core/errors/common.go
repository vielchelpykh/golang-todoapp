package core_errors

import "errors"

var (
	ErrNotFound        = errors.New("not found")
	ErrInvalidArgument = errors.New("invalid arguent")
	ErrConflict        = errors.New("conflict")
)
