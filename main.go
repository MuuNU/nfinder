package main

import (
	"log"
	"net/http"
	"nfinder/handle"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handle.HandlerMain)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
