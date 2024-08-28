package main

import (
	"encoding/json"
	"github.com/c00rni/rss-go/internal/database"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (cfg apiConfig) handleCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type request struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	inputData := request{}
	err := decoder.Decode(&inputData)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't decode parameters")
		return
	}

	today := time.Now()

	feed, err := cfg.DB.CreateFeeds(r.Context(), database.CreateFeedsParams{
		ID:        uuid.New(),
		CreatedAt: today,
		UpdatedAt: today,
		Name:      inputData.Name,
		Url:       inputData.Url,
		UserID:    user.ID,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create the feed")
		return
	}
	respondWithJSON(w, http.StatusCreated, feed)
}
