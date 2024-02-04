package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

// SearchFilms searches through a slice of films based on query parameters.
func SearchFilms(films []Film, query string) []Film {
    var results []Film
    for _, film := range films {
        if strings.Contains(strings.ToLower(film.Title), strings.ToLower(query)) ||
            strings.Contains(strings.ToLower(film.Director), strings.ToLower(query)){
            results = append(results, film)
        }
    }
    return results
}

// SearchHandler handles the search requests.
// Define the Film struct
// Remove the duplicate declaration of Film

// NewFilm creates a new Film instance.
func NewFilm(title, director string, year int, genre, synopsis string) Film {
	return Film{
		Title:    title,
		Director: director,
	
	}
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	// Example films slice. In a real app, this could be fetched from a database or global variable.
	films := []Film{
		NewFilm("The Godfather", "Francis Ford Coppola", 1972, "Crime, Drama", "The aging patriarch of an organized crime dynasty transfers control of his clandestine empire to his reluctant son."),
		// Add more films here
	}
	results := SearchFilms(films, query)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

// You would need to register this handler in your main.go
// http.HandleFunc("/search", SearchHandler)
