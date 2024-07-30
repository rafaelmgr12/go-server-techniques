package main

import (
	"io"
	"net/http"
	"testing"
)

var helloMsg = "Hello student!\n"

func TestAPISuccess(t *testing.T) {
	res, err := http.Get("http://localhost:8080/")
	if err != nil {
		t.Fatal(err)
	}

	msg, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	if helloMsg != string(msg) {
		t.Fatal("Unexpected message")
	}

}
