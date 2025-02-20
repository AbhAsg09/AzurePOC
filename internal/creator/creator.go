package creator

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Video struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Age   string
	City  string
	Phone string
}

func randomString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var sb strings.Builder
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		sb.WriteRune(letters[rng.Intn(len(letters))])
	}
	return sb.String()
}

func generateRandomVideo() Video {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	return Video{
		Name:  randomString(8),
		Age:   fmt.Sprintf("%d", rng.Intn(60)+18), // Random age between 18 and 77
		City:  randomString(10),                   // Random city name
		Phone: fmt.Sprintf("%d", rng.Intn(9000000000)+1000000000),
	}
}

func storeInDatabase(video Video, db *gorm.DB) {
	// Use a bulk insert operation to improve performance
	if err := db.Create(&video).Error; err != nil {
		log.Printf("Failed to store videos: %v", err)
		return
	}

	log.Printf("Successfully stored Data.")

}

func GenerateAndStoreVideos(db *gorm.DB) {
	for {
		video := generateRandomVideo()
		storeInDatabase(video, db)
	}
}
