package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initRouter() *mux.Router {
	r := mux.NewRouter()
	return r
}

// New initializes and returns a new Server.
func New(addr string) *Server {
	return &Server{
		HTTPServer: &http.Server{Addr: addr},
		Router:     initRouter(),
	}
}

// Shutdown stops the server.
func (srv *Server) Shutdown() error {
	return srv.HTTPServer.Shutdown(nil)
}

// ListenAndServe starts running the server and handling requests.
func (srv *Server) ListenAndServe() {
	log.Printf("Starting Server on %s", srv.HTTPServer.Addr)
	go func() {
		err := srv.HTTPServer.ListenAndServe()
		if err != nil {
			log.Fatalln(err)
		}
	}()
}
