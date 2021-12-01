package modules

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	proto "github.com/meysifk/movie_proto/movies-API"
)

func (m *Module) HandleMovies(w http.ResponseWriter, h *http.Request) {

	// get query string parameter for search and page
	q := h.URL.Query()
	searchString := q.Get("s")
	page := q.Get("page")

	if searchString == "" || page == "" {
		log.Println("Search word or page is required")
		res := GenericResponse{
			Success: false,
			Message: "search word or page is required",
		}
		toJSON(w, http.StatusOK, res)
		return
	}

	// convert string page to int
	p, err := strconv.Atoi(page)
	if err != nil {
		log.Println(err)
		res := GenericResponse{
			Success: false,
			Message: "page should numeric",
		}
		toJSON(w, http.StatusOK, res)
		return
	}

	movies, err := m.MovieClient.GetMovies(context.Background(), &proto.MovieParams{SearchWord: searchString, Page: int32(p)})
	if err != nil {
		log.Println(err)
		res := GenericResponse{
			Success: false,
			Message: "internal server error",
		}
		toJSON(w, http.StatusOK, res)
		return
	}

	res := GenericResponse{
		Success: true,
		Message: "success",
		Data:    movies,
	}

	toJSON(w, http.StatusOK, res)
}

func (m *Module) HandleMovie(w http.ResponseWriter, h *http.Request) {

	movieId := mux.Vars(h)["movie_id"]

	if movieId == "" {
		log.Println("movie id is required")
		res := GenericResponse{
			Success: false,
			Message: "movie id is required",
		}
		toJSON(w, http.StatusOK, res)
		return
	}

	movie, err := m.MovieClient.GetMovie(context.Background(), &proto.SingleMovieParams{ImdbId: movieId})
	if err != nil {
		log.Println(err)
		res := GenericResponse{
			Success: false,
			Message: "internal server error",
		}
		toJSON(w, http.StatusOK, res)
		return
	}

	res := GenericResponse{
		Success: true,
		Message: "success",
		Data:    movie,
	}

	toJSON(w, http.StatusOK, res)

}
