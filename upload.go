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
	if tzpe == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if tzpe != "url" && filename == "" {
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

func getToken(r *http.Request) string {
	cookie, err := r.Cookie("authToken")
	if err != nil {
		log.Println("getToken:", err)
		return ""
	}
	return cookie.Value
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
	userid, err := getUserIDFromJWT(getToken(r))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	uuid := generateUUID()
	database.CreateNewContent(&content{
		Id:     uuid,
		Data:   string(b),
		Ctype:  Link,
		UserId: userid,
	})
	w.WriteHeader(http.StatusOK)
	b, err = json.Marshal(uploadresponse{Uuid: uuid})
	_, _ = w.Write(b)
}

func fileupload(w http.ResponseWriter, r *http.Request) {
	tzpe := r.Header.Get("content-type")
	filename := r.Header.Get("x-filename")
	uuid := generateUUID()
	f, _ := os.Create("./cdn/userdata/" + uuid)
	defer f.Close()

	userid, err := getUserIDFromJWT(getToken(r))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		os.Remove(f.Name())
		return
	}

	_, err = io.Copy(f, r.Body)
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
		Id:     uuid,
		Ctype:  ctzpe,
		Data:   filename,
		UserId: userid,
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
