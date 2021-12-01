package modules

import (
	"context"
	"errors"
	"log"
)

func (m *Module) GetMovies(ctx context.Context, params *proto.MovieParams) (*proto.MoviesRes, error) {

	// validate request params
	if params.Page <= 0 {
		return nil, errors.New("page should bigger than zero")
	}

	if params.SearchWord == "" {
		return nil, errors.New("search word should not empty")
	}

	omdbMoviesParam := omdb.MoviesRequest{
		SearchWord: params.SearchWord,
		Page:       int(params.Page),
	}

	// request to omdb api
	omdbRes, resByte, err := m.OmdbApi.GetMovies(m.HttpClient, omdbMoviesParam)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// insert api response to log in DB
	err = m.InsertMovieLog(string(resByte))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// construct data for movie search list
	var movies []*proto.Movie

	for _, movie := range omdbRes.ResultSearch {
		m := proto.Movie{
			Title:  movie.Title,
			Year:   movie.Year,
			ImdbId: movie.ImdbID,
			Poster: movie.Poster,
		}

		movies = append(movies, &m)
	}

	res := proto.MoviesRes{
		Movies: movies,
	}

	return &res, nil
}

func (m *Module) GetMovie(ctx context.Context, params *proto.SingleMovieParams) (*proto.Movie, error) {

	// validate imdbId
	if params.ImdbId == "" {
		return nil, errors.New("imdbId should not empty")
	}

	// request data to movie api
	omdbMovie, resByte, err := m.OmdbApi.GetMovie(m.HttpClient, params.ImdbId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// insert api response to log in DB
	err = m.InsertMovieLog(string(resByte))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// construct data for movie
	movie := proto.Movie{
		Title:      omdbMovie.Title,
		Year:       omdbMovie.Year,
		Rated:      omdbMovie.Rated,
		Released:   omdbMovie.Released,
		Runtime:    omdbMovie.Runtime,
		Genre:      omdbMovie.Genre,
		Director:   omdbMovie.Director,
		Writer:     omdbMovie.Writer,
		Actors:     omdbMovie.Actors,
		Plot:       omdbMovie.Plot,
		Language:   omdbMovie.Language,
		Poster:     omdbMovie.Poster,
		Production: omdbMovie.Production,
		ImdbId:     omdbMovie.ImdbID,
		ImdbRating: omdbMovie.ImdbRating,
		ImdbVotes:  omdbMovie.ImdbVotes,
	}

	return &movie, nil
}
