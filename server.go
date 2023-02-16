package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"sync"

	"todolist/common"
)

func main() {
	db := common.Init()
	defer db.Close()

	srv := &http.Server{Addr: "localhost:8181"}
	httpServerExitDone := &sync.WaitGroup{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello world\n")
	})

	go func() {
		defer httpServerExitDone.Done() // let main know we are done cleaning up

		// always returns error. ErrServerClosed on graceful close
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			// unexpected error. port in use?
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	defer srv.Shutdown(context.TODO())
}
