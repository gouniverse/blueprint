package server

import (
	"context"
	"net/http"
	"testing"
	"time"
)

func TestServer_Start(t *testing.T) {
	srv := NewServer("localhost:45566", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello world"))
	})
	go func() {
		time.Sleep(1 * time.Second)
		srv.Shutdown(context.Background())
	}()
	err := srv.Start()
	if err != nil {
		t.Error("unexpected error:", err)
	}
}
