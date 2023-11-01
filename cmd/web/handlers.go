package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}

	w.Write([]byte("snippet box home"))
}

func snippetCreate(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "not allowed method", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("create snippet"))
}

func snippetView(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.Error(w, "invalid query parameter", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "show snippet with id = %d", id)
}
