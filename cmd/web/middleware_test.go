package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var myH myHandler
	h := NoSurf(&myH)

	switch v := h.(type) {
	case http.Handler:
		fmt.Print("It works")
	default:
		t.Error(fmt.Sprintf("Type is %t but exprected handler", v))
	}
}
