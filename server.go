package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hi! %s", req.URL.Path[1:])
	log.Println("Request: ", req)
}

func main() {
	port := os.Getenv("PORT")

	http.HandleFunc("/", handler)
	// log.Fatal(http.ListenAndServe(":8000", nil))
	log.Fatal(http.ListenAndServe(":"+port, nil))
}