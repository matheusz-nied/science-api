package service

import (
	"log"
	"time"
)

func StartCronJob(service APODService) {
	ticker := time.NewTicker(24 * time.Hour)
	go func() {
		for range ticker.C {
			log.Println("Running cron job daily...")
			if err := service.FetchAndSaveAPOD(); err != nil {
				log.Println("Error fetching and saving APoD:", err)
			}
		}
	}()
}
