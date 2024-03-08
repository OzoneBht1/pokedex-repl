package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func HandleFetch(url string) (ApiResponse, error) {
	res, err := http.Get(url)
	if err != nil {
		return ApiResponse{}, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ApiResponse{}, err
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return ApiResponse{}, err
	}
	response := ApiResponse{}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return ApiResponse{}, err
	}
	return response, nil
}
