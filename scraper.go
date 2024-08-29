package main

import (
	"context"
	"github.com/c00rni/rss-go/internal/database"
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
			wg.Add(1)
			go func(feed database.Feed) {
				_, err := fetchRSS(feed.Url)
				if err != nil {
					log.Printf("Couldn't fetch Feed at %s", feed.Url)
				} else {
					log.Printf("Feed %v updated.", feed.Name)
				}
				cfg.DB.MarkFeedFetched(context.Background(), feed.ID)
				defer wg.Done()
			}(feed)
		}
		wg.Wait()
	}
}
