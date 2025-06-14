package main

import (
	"fmt"
	"net/http"
	"time"

	"lightsaber.dkadev.xyz/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIdParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
	}
	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "SpiderMan",
		Runtime:   102,
		Genres:    []string{"drama", "action"},
		Version:   1,
	}

	err = app.writeJson(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
