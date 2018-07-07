package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/RyanTKing/wombats/pkg/server"
)

var (
	addr = flag.String("http", ":5080", "Server address")
)

func main() {
	srv := server.New(*addr)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	srv.ListenAndServe()

	<-c
	srv.Shutdown()
}
