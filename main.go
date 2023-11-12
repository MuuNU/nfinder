package main

import (
	"log"
	"net/http"
	"nfinder/handle"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/search", handle.HandlerSearch)
	http.Handle("/", http.RedirectHandler("/search", http.StatusSeeOther))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
