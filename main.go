package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	log.SetOutput(os.Stdout)

	log.Println("Starting up!")

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello there"))
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println(server.ListenAndServe())
}
