package main

import (
	_ "embed"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

//go:embed static/entry.html
var indexhtml []byte

var database *sqliteStore

func main() {
	var err error
	database, err = NewSqliteStore()
	if err != nil {
		log.Fatal(err)
	}
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./static"))
	mux.HandleFunc("GET /", loadReact)
	mux.HandleFunc("GET /{uuid}", loadReact)

	mux.HandleFunc("GET /api/{uuid}", lookupContent)
	//mux.HandleFunc("POST /api/upload")
	//mux.HandleFunc("PUT /api/{uuid}") //upload return'd welche uuid du uploaden kannst

	mux.Handle("GET /cdn/{filename}", http.StripPrefix("/cdn/", fs))

	err = http.ListenAndServe(":8888", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func lookupContent(w http.ResponseWriter, r *http.Request) {
	uuid := r.PathValue("uuid")
	log.Println(uuid)
	c, err := database.LookupContent(uuid)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	b, _ := json.Marshal(c)
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func loadReact(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write(indexhtml)
}

func generateUUID() string {
	f, err := os.Open("/proc/sys/kernel/random/uuid")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	b := make([]byte, 37)
	_, err = f.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}
