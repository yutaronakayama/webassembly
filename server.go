package main

import (
	"log"
	"net/http"
)

func server() {
	http.Handle("/", http.FileServer(http.Dir("")))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//server()
