package main

import (
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", List)
	http.ListenAndServe(":8080", r)
}

func List(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
  
	w.Header().Set("Content-Type", "application/json")
  
	if _, err := io.Copy(w, resp.Body); err != nil {
	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return
	}
}