package main

import (
	"log/slog"
	"net/http"
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
	app.clientError(w, http.StatusNotExtended)
}
