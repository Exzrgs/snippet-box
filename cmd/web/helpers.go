package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"bytes"
	"time"
)

func (app *application) serverError(w http.ResponseWriter, req *http.Request, err error) {
	var (
		method = req.Method
		url    = req.URL.RequestURI()
	)

	app.logger.Error(err.Error(), slog.String("method", method), slog.String("url", url))
	http.Error(w, "internal server error", http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) render(w http.ResponseWriter, req *http.Request, status int, page string, td templateData) {
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("template %s does not exist", page)
		app.serverError(w, req, err)
		return
	}

	buf := new(bytes.Buffer)
	err := ts.ExecuteTemplate(buf, "base", td)
	if err != nil {
		app.serverError(w, req, err)
		return
	}

	w.WriteHeader(status)

	err = ts.ExecuteTemplate(w, "base", td)
	if err != nil {
		app.serverError(w, req, err)
		return
	}
}

func (*application) newTemplateData()templateData{
	return templateData{
		CurrentYear: time.Now().Year(),
	}
}