package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(&payload)
	if err != nil {
		log.Println("Failed to convert payload into json")
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	w.Write(jsonData)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errorResponse{
		Error: msg,
	})
}

func extractAuthorization(r *http.Request, prefix string) (string, error) {
	target, found := strings.CutPrefix(r.Header.Get("Authorization"), prefix)
	if !found {
		return target, errors.New(fmt.Sprintf("Prefix '%v' not found.", prefix))
	}
	return target, nil
}

func fetchRSS(url string) (RSSFeed, error) {
	resp, err := http.Get(url)
	if err != nil {
		return RSSFeed{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return RSSFeed{}, err
	}

	rssFeed := RSSFeed{}
	err = xml.Unmarshal(body, &rssFeed)
	if err != nil {
		return RSSFeed{}, err
	}
	return rssFeed, nil
}
