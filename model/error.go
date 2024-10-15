package model

import "errors"

var (
	ErrNotFound     = errors.New("not found")
	ErrInvalidParam = errors.New("invalid param")
)
