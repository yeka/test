package main

import (
	"context"
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
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second)
		s := "Hello"
		if len(os.Args) > 1 {
			s = os.Args[1]
		}
		_, _ = fmt.Fprint(w, s)
	})
	s := http.Server{Addr: ":8123", Handler: mux}
	l, err := reuseport.Listen("tcp", s.Addr)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}

	go func() {
		err = s.Serve(l)
		if err != nil && err != http.ErrServerClosed {
			fmt.Printf("Error serving: %#v\n", err)
		}
	}()

	stopped := make(chan os.Signal)
	signal.Notify(stopped, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	<-stopped
	close(stopped)

	err = s.Shutdown(context.Background())
	if err != nil {
		log.Println("Shutdown error:", err)
		os.Exit(1)
	}
	log.Println("Shutdown gracefully")
}
