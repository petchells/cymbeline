package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type Turned struct {
	Turned []string
}

func serve() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		board := r.RequestURI
		log.Println(r.RequestURI)
		turned := Turned{[]string{"A1"}}
		json.NewEncoder(w).Encode(turned)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
