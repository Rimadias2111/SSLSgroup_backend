package emoji

import (
	database "backend/st_database"
	"context"
	"log"
	"time"
)

func StartEmojiUpdater(ctx context.Context, store *database.Store) {
	go func() {
		ticker := time.NewTicker(3 * time.Minute)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				log.Println("Stopping emoji updater")
				return
			case <-ticker.C:
				log.Println("Starting emoji update...")
				err := store.Logistic().Emoji(context.Background())
				if err != nil {
					log.Printf("Failed to update emojis: %v", err)
				} else {
					log.Println("Emoji update completed successfully")
				}
			}
		}
	}()
}
