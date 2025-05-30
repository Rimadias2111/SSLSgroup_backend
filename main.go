package main

import (
	"backend/api"
	"backend/api/controllers"
	emoji "backend/etc/emoji_updater"
	"backend/etc/search"
	"backend/models"
	"backend/service"
	database "backend/st_database"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

func setupDatabase() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=US/Eastern",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Connecting to DB with DSN: %s", dsn)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database %v", err)
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(30 * time.Minute)

	log.Println("Database connection established")

	if err := models.AutoMigrate(db); err != nil {
		log.Fatalf("Failed to auto migrate %v", err)
	}

	return db, nil
}

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	godotenv.Load(".env")
	db, err := setupDatabase()
	if err != nil {
		log.Fatalf("Failed to setup database %v", err)
	}

	store := database.New(db)
	serviceS := service.New(store)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	emoji.StartEmojiUpdater(ctx, store)

	cont := controllers.NewController(serviceS)

	errLoc := search.LoadLocations("/app/data/locations.json")
	if errLoc != nil {
		log.Fatalf("Ошибка загрузки данных: %v", errLoc)
	}

	router := api.Construct(*cont)
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
