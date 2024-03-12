package service

import "errors"

var (
	ErrUniqueViolation = errors.New("record already exists")
	ErrNotFound        = errors.New("nothing found")
)
