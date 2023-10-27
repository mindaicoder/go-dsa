package linkedlist

import "errors"

var (
	ErrEmpty           = errors.New("list is empty")
	ErrNotEmpty        = errors.New("list is not empty")
	ErrIndexOutOfRange = errors.New("index out of range")
)
