package main

import "testing"

func failIf(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}
