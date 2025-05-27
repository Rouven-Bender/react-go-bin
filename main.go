package main

import (
	_ "embed"
	"log"
	"net/http"
)

//go:embed static/entry.html
var indexhtml []byte

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./static"))
	mux.HandleFunc("GET /", homepage)

	mux.Handle("GET /cdn/{filename}", http.StripPrefix("/cdn/", fs))

	err := http.ListenAndServe(":8888", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func homepage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write(indexhtml)
}
