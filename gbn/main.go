package main

import (
	//"fmt"
	"log"
	"net/http"
)

const (
	gitURL    = "https://github.com"
	gitAPIURL = "https://api.github.com"
	gbnURL    = "/notifications/mail/"
	gbnPort   = ":10005"
)

func init() {
	log.SetFlags(0)
}

func main() {
	http.HandleFunc(gbnURL, handler)
	http.ListenAndServe(gbnPort, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
}
