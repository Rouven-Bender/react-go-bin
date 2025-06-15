package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

//go:embed cdn/entry.html
var indexhtml []byte

var database *sqliteStore

func main() {
	if len(os.Args) < 2 {
		fmt.Println("expected subcommand")
		fmt.Printf("list of subcommands:\n\t- serve\n")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "serve":
		initDatabase()
		launchServer()
	case "add-user":
		initDatabase()
		addUser()
	}
}

func addUser() {
	newUserSecret := makeSecretKey()
	hash := hashSecretKey(newUserSecret)
	err := database.AddUser(hash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Secret Key for new User: %s", newUserSecret)
}

func initDatabase() {
	var err error
	database, err = NewSqliteStore()
	if err != nil {
		log.Fatal(err)
	}
}

func launchServer() {
	mux := http.NewServeMux()
	static := http.FileServer(http.Dir("./cdn"))
	userdata := http.FileServer(http.Dir("./cdn/userdata/"))

	mux.Handle("GET /cdn/{filename}", http.StripPrefix("/cdn/", static))
	mux.Handle("GET /userdata/{filename}", http.StripPrefix("/userdata/", userdata))

	mux.HandleFunc("GET /", loadReact)
	mux.HandleFunc("GET /{uuid}", loadReact)
	mux.HandleFunc("GET /login", loadReact)
	mux.HandleFunc("GET /upload", loadReact)

	mux.HandleFunc("GET /api/lookup/{uuid}", lookupContent)
	mux.HandleFunc("POST /api/login", verifyCred)
	mux.HandleFunc("POST /api/upload", requiresAuthToken(uploadHandler))

	log.Println("Server launching on port 8888")
	err := http.ListenAndServe(":8888", mux)
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
	b = bytes.Trim(b, "\n")
	return string(b)
}
