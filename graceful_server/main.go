package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
)

var port = "8080"

type gracefulServer struct {
	httpServer *http.Server
	listener   net.Listener
}

func (server *gracefulServer) start() error {
	listener, err := net.Listen("tcp", server.httpServer.Addr)
	if err != nil {
		return err
	}

	server.listener = listener
	go server.httpServer.Serve(server.listener)
	fmt.Println("Server now listening on " + server.httpServer.Addr)
	return nil
}

func handle(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello student!\n")
}

// Initialize a new server instance and return reference.
func newServer(port string) *gracefulServer {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handle)
	httpServer := &http.Server{Addr: ":" + port, Handler: mux}
	return &gracefulServer{httpServer: httpServer}
}

func main() {
	server := newServer(port)
	server.start()
}
