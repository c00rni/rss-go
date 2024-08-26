package main

import (
	"net/http"
)

func handleHealthz(w http.ResponseWriter, req *http.Request) {
	type response struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, http.StatusOK, response{Error: "ok"})
	return
}
