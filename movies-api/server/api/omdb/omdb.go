package omdb

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/meysifk/test-backend/movie-API/server/network"
)

func (o *Omdb) GetMovies(httpClient *http.Client, params MoviesRequest) (*MoviesResponse, []byte, error) {

	// query params
	queryParams := map[string]string{
		"apikey": o.Credential,
		"s":      params.SearchWord,
		"page":   fmt.Sprintf("%d", params.Page),
	}

	resByte, err := network.Do(httpClient, http.MethodGet, o.Host, queryParams, nil, nil)
	if err != nil {
		log.Print(err)
		return nil, nil, err
	}

	var res MoviesResponse

	err = json.Unmarshal(resByte, &res)
	if err != nil {
		log.Print(err)
		return nil, nil, err
	}

	return &res, resByte, nil

}

func (o *Omdb) GetMovie(httpClient *http.Client, imdbId string) (*MovieResponse, []byte, error) {

	// query params
	queryParams := map[string]string{
		"apikey": o.Credential,
		"i":      imdbId,
	}

	resByte, err := network.Do(httpClient, http.MethodGet, o.Host, queryParams, nil, nil)
	if err != nil {
		log.Print(err)
		return nil, nil, err
	}

	var res MovieResponse

	err = json.Unmarshal(resByte, &res)
	if err != nil {
		log.Print(err)
		return nil, nil, err
	}

	return &res, resByte, nil

}
