package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"snippet-box/internal/models"

	"github.com/julienschmidt/httprouter"
)

func (app *application) home(w http.ResponseWriter, req *http.Request) {
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

func (app *application) snippetCreateView(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Display the form for creating a new snippet..."))
}

func (app *application) snippetCreate(w http.ResponseWriter, req *http.Request) {
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n- Kobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, req, err)
		return
	}

	http.Redirect(w, req, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}

func (app *application) snippetView(w http.ResponseWriter, req *http.Request) {
	params := httprouter.ParamsFromContext(req.Context())

	id, err := strconv.Atoi(params.ByName("id"))
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
