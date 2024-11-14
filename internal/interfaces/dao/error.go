package dao

import "errors"

var (
	ErrNoUpdateHappened = errors.New("no affected rows when trying to update")
	ErrNilParam         = errors.New("param cannot be nil")
	ErrDuplicate        = errors.New("duplicate entry")
	ErrNoResult         = errors.New("no result")
)
