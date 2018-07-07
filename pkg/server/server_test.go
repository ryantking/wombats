package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testPort = ":5080"

func TestNewServer(t *testing.T) {
	srv := New(testPort)
	assert.NotNil(t, srv)
	assert.Equal(t, testPort, srv.HTTPServer.Addr)
}

func TestListenAndServeThenShutdown(t *testing.T) {
	srv := New(testPort)
	assert.NotNil(t, srv)
	srv.ListenAndServe()
	err := srv.Shutdown()
	assert.Nil(t, err)
}
