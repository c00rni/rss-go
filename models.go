package main

import (
	"github.com/c00rni/rss-go/internal/database"
	"github.com/google/uuid"
	"time"
)

type Feed struct {
	ID            uuid.UUID `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Name          string    `json:"name"`
	Url           string    `json:"url"`
	UserID        uuid.UUID `json:"user_id"`
	LastFetchedAt time.Time `json:"last_fetched_at"`
}

func databaseFeedToFeed(feed database.Feed) Feed {
	lastFetchedAt := feed.LastFetchedAt.Time
	if !feed.LastFetchedAt.Valid {
		lastFetchedAt = time.Unix(0, 0)
	}
	return Feed{
		ID:            feed.ID,
		CreatedAt:     feed.CreatedAt,
		UpdatedAt:     feed.UpdatedAt,
		Name:          feed.Name,
		Url:           feed.Url,
		UserID:        feed.UserID,
		LastFetchedAt: lastFetchedAt,
	}
}
