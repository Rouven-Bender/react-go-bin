package main

import (
	"bytes"
	"crypto/rand"
	_ "embed"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/argon2"
)

//go:embed JWT_SECRET
var JWT_SECRET []byte

//go:embed SALT
var salt []byte

const issuerName = "personal-pastebin"

type loginrequest struct {
	Key string `json:"key"`
}

type loginresponse struct {
	Token   string `json:"token,omitzero"`
	Message string `json:"msg,omitzero"`
}

func makeSecretKey() string {
	return rand.Text()
}

// hash the secret key and return base64 string
func hashSecretKey(key string) string {
	return base64.StdEncoding.EncodeToString(argon2.IDKey([]byte(key), salt, 1, 64*1024, 4, 32))
}

func verifyCred(w http.ResponseWriter, r *http.Request) {
	if t := r.Header.Get("Content-type"); t != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	buf := make([]byte, 1024)
	_, err := r.Body.Read(buf)
	if !errors.Is(err, io.EOF) {
		log.Println("error reading the request body:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	buf = bytes.Trim(buf, "\x00") // trim the null terminator the browser sends
	req := loginrequest{}
	err = json.Unmarshal(buf, &req)
	if err != nil {
		log.Println("error unmarshaling json:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	hash := hashSecretKey(req.Key)
	if userid, ok := database.UserIdForSecretKey(hash); ok {
		if userid == -1 {
			log.Fatalf("negativ 1 gotten for %s as hash %s", req.Key, hash)
		}
		token, err := createJWTToken(userid)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("ERROR: ", err)
			return
		}
		b, err := json.Marshal(loginresponse{Token: token})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("ERROR: ", err)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(b)
		return
	} else {
		b, err := json.Marshal(loginresponse{Message: "wrong credentials"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(b)
		return
	}
}
func createJWTToken(userid int) (string, error) {
	claims := &jwt.RegisteredClaims{
		ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Hour * 2)},
		Issuer:    issuerName,
		Subject:   strconv.Itoa(userid),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JWT_SECRET)
	if err != nil {
		return "", fmt.Errorf("signing token: %w", err)
	}
	return tokenString, nil
}

func requiresAuthToken(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("authToken")
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusForbidden)
			return
		}
		token, err := validateJWT(cookie.Value)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusForbidden)
			return
		}
		if !token.Valid {
			log.Println(err)
			w.WriteHeader(http.StatusForbidden)
			return
		}
		f(w, r)
	}
}

func getUserIDFromJWT(tokenstring string) (int, error) {
	token, err := validateJWT(tokenstring)
	if err != nil {
		return -1, err
	}
	s, err := token.Claims.GetSubject()
	if err != nil {
		return -1, err
	}
	id, err := strconv.Atoi(s)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func validateJWT(tokenstring string) (*jwt.Token, error) {
	return jwt.Parse(tokenstring,
		func(token *jwt.Token) (any, error) {
			if token.Method.Alg() == "HS256" {
				if issuer, err := token.Claims.GetIssuer(); err == nil && issuer == issuerName {
					return JWT_SECRET, nil
				}
			}
			return nil, fmt.Errorf("Couldn't Parse Token")
		},
	)
}
