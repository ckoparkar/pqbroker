package main

import (
	"errors"
	"net/http"
	"regexp"
	"testing"
)

type ErrorWithCode struct {
	Err  error
	Code int
}

var (
	// status 409
	ErrInstanceAlreadyExists  = ErrorWithCode{Err: errors.New("instance already exists"), Code: http.StatusConflict}
	RInstanceAlreadyExists, _ = regexp.Compile(".*database \".*\" already exists")

	// status 500
	ErrServerNotReachable = ErrorWithCode{Err: errors.New("server not reachable"), Code: http.StatusInternalServerError}
)

func failIf(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func pqError(err error) *ErrorWithCode {
	e := err.Error()
	switch {
	case RInstanceAlreadyExists.MatchString(e):
		return &ErrInstanceAlreadyExists
	}
	return &ErrorWithCode{err, 0}
}
