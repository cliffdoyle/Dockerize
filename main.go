package main

import (
	"fmt"
	"log"
	"net/http"

	backend "ascii-art/Backend"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", backend.Webhandler)
	mux.HandleFunc("/submit", backend.HandleRequest)
	fmt.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
