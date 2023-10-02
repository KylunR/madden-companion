package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, req *http.Request) {

	reqBody, err := io.ReadAll(req.Body)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Request Body: %s", reqBody)

	fmt.Fprintf(w, "Hi! %s", req.URL.Path[1:])
	// log.Println("Request: ", req)

	w.WriteHeader(http.StatusOK)
}

func main() {
	port := os.Getenv("PORT")

	http.HandleFunc("/pc/3334523/leagueteams", handler)
	// log.Fatal(http.ListenAndServe(":8000", nil))
	log.Fatal(http.ListenAndServe(":"+port, nil))
}