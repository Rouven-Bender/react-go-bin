package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
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
	case "url":
		writeLinkToDB(w, r)
	default:
		fileupload(w, r)
	}
}

type uploadresponse struct {
	Uuid string `json:"uuid"`
}

func writeLinkToDB(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 2*1024)
	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	database.CreateNewContent(&content{
		Id:    generateUUID(),
		Data:  string(b),
		Ctype: Link,
	})
}

func fileupload(w http.ResponseWriter, r *http.Request) {
	tzpe := r.Header.Get("content-type")
	filename := r.Header.Get("x-filename")
	uuid := generateUUID()
	f, _ := os.Create("./cdn/userdata/" + uuid)
	defer f.Close()

	_, err := io.Copy(f, r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		os.Remove(f.Name())
		return
	}

	var ctzpe contenttype

	switch strings.Split(tzpe, "/")[0] {
	case "image":
		ctzpe = Image
	default:
		ctzpe = Plaintext // TODO: make bin files there own category
	}

	err = database.CreateNewContent(&content{
		Id:    uuid,
		Ctype: ctzpe,
		Data:  filename,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("database write:", err)
		return
	}
	b, err := json.Marshal(uploadresponse{Uuid: uuid})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("fileupload:", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(b)
	if err != nil {
		log.Println(err)
	}
}
