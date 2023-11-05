package main

import (
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}

	/*
		実行している場所からの相対パスだからやばい
	*/
	files := []string{
		"./ui/html/pages/home.html",
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.logger.Error(err.Error(), slog.String("method", req.Method), slog.String("url", req.URL.RequestURI()))
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.logger.Error(err.Error(), slog.String("method", req.Method), slog.String("url", req.URL.RequestURI()))
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func (app *application) snippetCreate(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "not allowed method", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("create snippet"))
}

func (app *application) snippetView(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.Error(w, "invalid query parameter", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "show snippet with id = %d", id)
}
