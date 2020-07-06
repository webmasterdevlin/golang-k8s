package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", indexHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

type Object map[string]interface{ }

func indexHandler(w http.ResponseWriter, r *http.Request) {

	data := Object{
		"id":"9823hoh",
		"username": "this is GoLang",
	}
	_ = json.NewEncoder(w).Encode(data)
}
