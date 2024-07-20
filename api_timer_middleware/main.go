package main

/**
Please add new middleware that will measure the time taken by our APIs to respond in
milliseconds and print the results to logs.

**/

import (
	"flag"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var welcomeMsg = "Welcome to the graceful server! 💃🏼\n"

type gracefulServer struct {
	httpServer *http.Server
	listener   net.Listener
}

func withSimpleLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Default().Printf("Incoming traffic on route %s", r.URL.Path)
		handler.ServeHTTP(w, r)
	})
}

func withExecutionTime(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		handler.ServeHTTP(w, r)
		defer log.Default().Printf("Execution time - %d µs", time.Since(startTime).Milliseconds())
	})

}

func (server *gracefulServer) preStart() {
	server.httpServer.Handler = withExecutionTime(withSimpleLogger(server.httpServer.Handler)) // chain the middleware
}

func (server *gracefulServer) start() error {
	listener, err := net.Listen(
		"tcp",
		server.httpServer.Addr)
	if err != nil {
		return err
	}

	server.listener = listener
	go server.httpServer.Serve(server.listener)
	log.Default().Printf("Server now listening on %s", server.httpServer.Addr)
	return nil
}

func (s *gracefulServer) shutdown() error {
	if s.listener != nil {
		err := s.listener.Close()
		s.listener = nil
		if err != nil {
			return err
		}
	}

	log.Default().Println("Shutting down server")
	return nil
}

func handle(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello student!\n")
}

func greetingHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, welcomeMsg)
}

func newServer(port string) *gracefulServer {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handle)
	mux.HandleFunc("/greeting", greetingHandler)

	httpServer := &http.Server{Addr: ":" + port, Handler: mux}
	return &gracefulServer{httpServer: httpServer}
}

func main() {
	var port string
	flag.StringVar(&port, "port", "8080", "./course_server -port 8080")
	flag.Parse()

	done := make(chan bool, 1)
	interrupts := make(chan os.Signal, 1)
	signal.Notify(interrupts, syscall.SIGINT, syscall.SIGTERM)

	server := newServer(port)
	server.preStart()
	err := server.start()
	if err != nil {
		log.Fatalf("Error starting server - %v\n", err)
	}

	go func() {
		sig := <-interrupts
		log.Default().Printf("Signal intercepted - %v\n", sig)
		server.shutdown()
		done <- true
	}()

	<-done
}
