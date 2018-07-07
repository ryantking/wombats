package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Server is a daemon server that handles requests for packages
type Server struct {
	HTTPServer *http.Server
	Router     *mux.Router
}
