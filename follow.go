package main

import (
	"encoding/json"
	"fmt"
	"github.com/c00rni/rss-go/internal/database"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

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

func (cfg apiConfig) handleUnfollowFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowID := r.PathValue("feedFollowID")
	if feedFollowID == "" {
		respondWithError(w, http.StatusBadRequest, "The feed id to delete must be set.")
		return
	}

	feedID, err := uuid.Parse(feedFollowID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v is not a valid UUID", feedFollowID))
		return
	}

	err = cfg.DB.Unfollow(r.Context(), database.UnfollowParams{
		FeedID: feedID,
		UserID: user.ID,
	})

	if err != nil {
		log.Printf("Failed to unfollow (User: %v, Feed: %v  -  %v", user.ID, feedID, err)
		respondWithError(w, http.StatusBadRequest, "Can't unfollow the feed.")
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (cfg apiConfig) handleGetUserFeeds(w http.ResponseWriter, r *http.Request, user database.User) {
	feeds, err := cfg.DB.GetUserFeeds(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Unauthorized user.")
		return
	}
	respondWithJSON(w, http.StatusOK, feeds)
}
