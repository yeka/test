package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	reuseport "github.com/libp2p/go-reuseport"
)

func main() {
	tctx := TerminationContext()
	server := NewServer()
	err := ListenAndServe(tctx, server)
	if err != nil {
		log.Println("Error:", err)
	} else {
		log.Println("Graceful shutdown")
	}
}

func NewServer() *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second)
		s := "Hello"
		if len(os.Args) > 1 {
			s = os.Args[1]
		}
		_, _ = fmt.Fprint(w, s)
	})
	return &http.Server{Addr: ":8123", Handler: mux}
}

func ListenAndServe(tctx context.Context, s *http.Server) error {
	l, err := reuseport.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}

	done := make(chan error)
	defer close(done)
	go func() {
		if err := s.Serve(l); err != nil && err != http.ErrServerClosed {
			done <- errors.New("serve error: " + err.Error())
		}
	}()

	select {
	case err := <-done:
		return err
	case <-tctx.Done():
		err := s.Shutdown(context.Background())
		if err != nil {
			log.Println("Shutdown error:", err)
		}
	}

	return nil
}

func TerminationContext() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		stopped := make(chan os.Signal)
		signal.Notify(stopped, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
		<-stopped
		cancel()
	}()
	return ctx
}
