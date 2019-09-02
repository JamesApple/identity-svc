package root

import (
	"errors"
	"fmt"
)

const (
	ErrNotFound AppErrorType = iota
	ErrUnexpected
	ErrInvalidInput
	ErrNotAvailable
)

type AppErrorType int

type AppError interface {
	Error() string
	Type() AppErrorType
}

type err struct {
	error
	t AppErrorType
}

func (e err) Type() AppErrorType {
	return e.t
}

func ErrWrap(t AppErrorType, e error) error {
	return &err{error: e, t: t}
}

func Errf(t AppErrorType, cause string, values ...interface{}) error {
	return ErrWrap(t, fmt.Errorf(cause, values...))
}

func Err(t AppErrorType, cause string) error {
	return ErrWrap(t, errors.New(cause))
}
