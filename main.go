package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	log.SetOutput(os.Stdout)

	log.Println("Starting up!")

	mux := http.NewServeMux()

	content, _ := ioutil.ReadFile("/etc/assets/stuff.txt")

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello there \n"))
		w.Write(content)
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println(server.ListenAndServe())
}
