package main

import (
	"fmt"
   "log"
	"io"
   "html/template"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
   w.Header().Add("Server", "Go")

   files := []string {
      "./ui/html/base.tmpl",
      "./ui/html/partials/nav.tmpl",
      "./ui/html/pages/home.tmpl",
   }

   ts, err := template.ParseFiles(files...)
   if err != nil {
      log.Print(err.Error())
      http.Error(w, "Internal Server Error", http.StatusInternalServerError)
      return
   }

   err = ts.ExecuteTemplate(w, "base", nil)
   if err != nil {
      log.Print(err.Error())
      http.Error(w, "Internal Server Error", http.StatusInternalServerError)
   }
}

func snippetView(w http.ResponseWriter, r *http.Request) {
   w.Header().Add("Server", "Go")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
   fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
   w.Header().Add("Server", "Go")
	io.WriteString(w, "Display a form for creating a new snippet...")
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
   w.Header().Add("Server", "Go")
   w.WriteHeader(http.StatusCreated)
   io.WriteString(w, "Save a new snippet...")
}
