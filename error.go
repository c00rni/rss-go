package main

import "net/http"

func handleError(w http.ResponseWriter, _ *http.Request) {
	type response struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, http.StatusInternalServerError, response{Error: "Internal Server Error"})
	return
}
