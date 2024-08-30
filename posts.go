package main

import (
	"github.com/c00rni/rss-go/internal/database"
	"log"
	"net/http"
)

func (cfg apiConfig) handleGetUserPosts(w http.ResponseWriter, r *http.Request, user database.User) {
	databasePosts, err := cfg.DB.GetUserPosts(r.Context(), user.ID)
	if err != nil {
		log.Printf("Failed to gets user posts : %v", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to get user posts")
	}

	posts := []Post{}
	for _, post := range databasePosts {
		posts = append(posts, databasePostToPost(post))
	}

	respondWithJSON(w, http.StatusOK, posts)
}
