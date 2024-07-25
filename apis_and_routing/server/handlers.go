package server

import (
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var helloMsg = "Hello student!\n"
var welcomeMsg = "Welcome to the graceful server! 💃🏼\n"

func baseHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, helloMsg)
}

func greetingHandler(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	greeting := welcomeMsg
	if name != "" {
		greeting = "Hello " + name + "!\n" + welcomeMsg
	} else {
		name = r.URL.Query().Get("name")
		if name != "" {
			greeting = "Hello " + name + "!\n" + welcomeMsg
		}
	}

	io.WriteString(w, greeting)
}

// ADD HANDLER HERE
