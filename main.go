// TaskSystem - CTCL 2023-2025
// File: main.go
// Purpose:
// Created: March 14, 2025
// Modified: March 14, 2025

package main

import (
	//"github.com/flosch/pongo2/v6"
	//"github.com/emersion/go-webdav"
	"net/http"
)

func route(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func main() {
	http.HandleFunc("/", route)
	// TODO: have port defined by config
	s := &http.Server{
		Addr: ":8080",
	}
	s.ListenAndServe()
}
