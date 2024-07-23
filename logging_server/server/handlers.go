package server

import (
	"io"
	"net/http"
)

var helloMsg = "Hello student!\n"
var welcomeMsg = "Welcome to the graceful server! 💃🏼\n"

func handle(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, helloMsg)
}

func greetingHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, welcomeMsg)
}
