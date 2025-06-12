package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 10*1024*1024)
	tzpe := r.Header.Get("content-type")
	filename := r.Header.Get("x-filename")
	if tzpe == "" || filename == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	switch tzpe {
	case "image/png":
	}
	uuid := generateUUID()
	f, _ := os.Create("./cdn/userdata/" + uuid)
	io.Copy(f, r.Body)
	log.Println(uuid)
}
