package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var t interface{}
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	log.Println("Request body: ", t)

	fmt.Fprintf(w, "Hi! %s", req.URL.Path[1:])
	// log.Println("Request: ", req)

	w.WriteHeader(http.StatusOK)
}

func main() {
	port := os.Getenv("PORT")

	http.HandleFunc("/", handler)
	// log.Fatal(http.ListenAndServe(":8000", nil))
	log.Fatal(http.ListenAndServe(":"+port, nil))
}