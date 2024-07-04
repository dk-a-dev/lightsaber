package main

import (
	"fmt"
	"net/http"
	"time"

	"lightsabor.dkadev.net/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Creates a Movie Handler")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIdParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Yeh Jawani hai Dewaani",
		Runtime:   120,
		Genres:    []string{"drama", "comedy", "romance"},
		Version:   1,
	}

	err = app.writeJson(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
