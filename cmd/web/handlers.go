package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"snippet-box/internal/models"
)

func (app *application) home(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		app.notFound(w)
		return
	}

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, req, err)
		return
	}
	td := app.newTemplateData()
	td.Snippets = snippets
	page := "home.html"
	app.render(w, req, http.StatusOK, page, td)
}

func (app *application) snippetCreate(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n- Kobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, req, err)
		return
	}

	http.Redirect(w, req, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
}

func (app *application) snippetView(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, req, err)
		}
		return
	}
	td := app.newTemplateData()
	td.Snippet = snippet
	page := "view.html"
	app.render(w, req, http.StatusOK, page, td)
}
