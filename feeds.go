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

	type respond struct {
		Feed        Feed                  `json:"feed"`
		Feed_Follow database.Feedfollowed `json:"feed_followed"`
	}

	decoder := json.NewDecoder(r.Body)
	inputData := request{}
	err := decoder.Decode(&inputData)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't decode parameters")
		return
	}

	today := time.Now()
	feedId := uuid.New()

	databaseFeed, err := cfg.DB.CreateFeeds(r.Context(), database.CreateFeedsParams{
		ID:        feedId,
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

	// Automaticly follow the feed
	feedFollow, _ := cfg.DB.FollowFeed(r.Context(), database.FollowFeedParams{
		FeedID:    feedId,
		UserID:    user.ID,
		CreatedAt: today,
		UpdatedAt: today,
	})

	respondWithJSON(w, http.StatusCreated, respond{Feed: databaseFeedToFeed(databaseFeed), Feed_Follow: feedFollow})
}

func (cfg apiConfig) handleGetFeeds(w http.ResponseWriter, r *http.Request) {
	databaseFeeds, err := cfg.DB.GetFeeds(r.Context())
	if err != nil {
		log.Printf("Failed to fetch data from the DB : %s", err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't get the feeds.")
		return
	}

	feeds := []Feed{}
	for _, feed := range databaseFeeds {
		feeds = append(feeds, databaseFeedToFeed(feed))
	}

	respondWithJSON(w, http.StatusCreated, feeds)
}
