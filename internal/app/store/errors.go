package store

import "errors"

var (
	ErrStudentNotFound = errors.New("Student not found")
	ErrWrongRequset = errors.New("Wrong request")
	InternalError = errors.New("Internal error")
)
