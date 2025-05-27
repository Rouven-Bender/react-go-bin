package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./static"))
	//mux.HandleFunc("GET /", homepage)

	mux.Handle("GET /cdn/{filename}", http.StripPrefix("/cdn/", fs))

	err := http.ListenAndServe(":8888", mux)
	if err != nil {
		log.Fatal(err)
	}
}
