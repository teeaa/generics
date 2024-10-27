package internal

import "errors"

var (
	ErrNotFound    = errors.New("not found")
	ErrServerError = errors.New("server error")
)
