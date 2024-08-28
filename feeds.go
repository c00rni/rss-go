package main

import (
	"encoding/json"
	"github.com/c00rni/rss-go/internal/database"
	"github.com/google/uuid"
	"log"
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
		log.Printf("Failed to create ther user %s : %s", inputData.Name, err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't create the feed")
		return
	}
	respondWithJSON(w, http.StatusCreated, feed)
}

func (cfg apiConfig) handleGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := cfg.DB.GetFeeds(r.Context())
	if err != nil {
		log.Printf("Failed to fetch data from the DB : %s", err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't get the feeds.")
		return
	}
	respondWithJSON(w, http.StatusCreated, feeds)
}

func (cfg apiConfig) handleFollowingFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type request struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	inputData := request{}
	err := decoder.Decode(&inputData)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't decode parameters")
		return
	}

	today := time.Now()

	follow, err := cfg.DB.FollowFeed(r.Context(), database.FollowFeedParams{
		FeedID:    inputData.FeedID,
		UserID:    user.ID,
		CreatedAt: today,
		UpdatedAt: today,
	})

	if err != nil {
		log.Printf("Failed to follow the feed %s : %s", inputData.FeedID, err)
		respondWithError(w, http.StatusInternalServerError, "Can't follow the feed.")
		return
	}
	respondWithJSON(w, http.StatusCreated, follow)

}
