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

type Post struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Url         string    `json:"url"`
	Description string    `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	FeedID      uuid.UUID `json:"feed_id"`
}

func databasePostToPost(post database.Post) Post {
	return Post{
		ID:          post.ID,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
		Title:       post.Title,
		Url:         post.Url,
		Description: post.Description.String,
		FeedID:      post.FeedID,
		PublishedAt: post.PublishedAt,
	}
}

// RSS FEED

type RSSFeed struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Title       string    `xml:"title"`
	Link        string    `xml:"link"`
	Description string    `xml:"description"`
	Language    string    `xml:"language"`
	Items       []RSSItem `xml:"item"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	PubDate     string `xml:"pubDate"`
	Guid        string `xml:"guid"`
	Description string `xml:"description"`
}
