package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello from Snippetbox")
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
   fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Display a form for creating a new snippet...")
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
   w.Header().Add("Server", "Go")
   w.WriteHeader(http.StatusCreated)
   io.WriteString(w, "Save a new snippet...")
}
