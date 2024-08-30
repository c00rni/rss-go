package main

import (
	"context"
	"database/sql"
	"github.com/c00rni/rss-go/internal/database"
	"github.com/google/uuid"
	"log"
	"sync"
	"time"
)

func (cfg apiConfig) scrapFeeds(timeBetweenRequest time.Duration, limit int) {
	ticker := time.NewTicker(timeBetweenRequest)
	var wg sync.WaitGroup
	for ; ; <-ticker.C {

		feeds, err := cfg.DB.GetNextFeedsToFetch(context.Background(), int32(limit))
		if err != nil {
			log.Println(err)
			continue
		}

		for _, feed := range feeds {
			rssFeed, err := fetchRSS(feed.Url)
			if err != nil {
				log.Printf("Couldn't fetch Feed at %s", feed.Url)
				continue
			}
			for _, postItem := range rssFeed.Channel.Items {
				wg.Add(1)
				go func() {

					postId, err := uuid.NewUUID()
					if err != nil {
						log.Printf("Coudn't create post uuid : %v", err)
						return
					}

					description := sql.NullString{String: postItem.Description}
					if postItem.Description != "" {
						description.Valid = true
					} else {
						description.Valid = false
					}

					today := time.Now()

					publishedTime, err := time.Parse(time.RFC3339, postItem.PubDate)
					if err != nil {
						publishedTime = today
					}

					_, err = cfg.DB.CreatePost(context.Background(), database.CreatePostParams{
						ID:          postId,
						CreatedAt:   today,
						UpdatedAt:   today,
						Title:       postItem.Title,
						Url:         postItem.Link,
						Description: description,
						PublishedAt: publishedTime,
						FeedID:      feed.ID,
					})
					if err != nil {
						return
					}

					log.Printf("[New Post] - Feed %s - Title: %s", feed.Name, postItem.Title)
					defer wg.Done()
				}()
			}
			cfg.DB.MarkFeedFetched(context.Background(), feed.ID)
			log.Printf("Feed %v updated.", feed.Name)
		}
		wg.Wait()
	}
}
