package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	quitChan := make(chan os.Signal)
	mux.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		quitChan <- syscall.SIGINT
	})

	server := http.Server{
		Handler: mux,
		Addr:    ":8000",
	}

	g.Go(func() error {
		fmt.Println("server listenning")
		return server.ListenAndServe()
	})

	g.Go(func() error {
		select {
		case <-ctx.Done():
			fmt.Println("quit")
		case <-quitChan:
			fmt.Println("server shutdown")
		}
		timeoutCtx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		return server.Shutdown(timeoutCtx)
	})

	g.Go(func() error {
		notifySigns := []os.Signal{syscall.SIGINT, syscall.SIGTERM}
		signs := make(chan os.Signal)
		signal.Notify(signs, notifySigns...)

		select {
		case <-ctx.Done():
			return ctx.Err()
		case sig := <-signs:
			return fmt.Errorf("os signal: %#v", sig)
		}
	})
	fmt.Printf("quit: %+v\n", g.Wait())
}
