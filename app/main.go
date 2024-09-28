package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	r := mux.NewRouter()
	r = routes(r)
	ctx := context.Background()
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	sig := <-c

	fmt.Printf("Got interrupt signal, aborting... %v\n", sig)
	timeoutCtx, cancel := context.WithTimeout(ctx, 1*time.Minute)
	if err := server.Shutdown(timeoutCtx); err != nil {
		panic(err)
	}
	cancel()
	os.Exit(0)
}
