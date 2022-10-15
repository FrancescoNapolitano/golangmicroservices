package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type HelloHandler struct {
	l *log.Logger
}

type Person struct {
	Name     string
	Age      int
	UserInfo UserInfo
}

type UserInfo struct {
	Username string
	Email    string
}

func NewHelloHandler(l *log.Logger) *HelloHandler {
	return &HelloHandler{l}
}

func (h *HelloHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var p Person
	err := json.NewDecoder(r.Body).Decode(&p)
	fmt.Printf("Hello %s", p.Name)
	if err != nil {
		http.Error(rw, "Error occured", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(rw, "Person: %+v", p)
}
