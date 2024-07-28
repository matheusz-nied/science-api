package main

import (
	"encoding/json"
	"log"
	db "nied-science/internal/database"
	"nied-science/internal/model"
	"nied-science/internal/repository"
	"os"
)

func main() {
	db.Init()

	data, err := os.ReadFile("scripts/apods.json")
	if err != nil {
		log.Fatalf("failed to read JSON file: %v", err)
	}

	var apods []model.APOD
	if err := json.Unmarshal(data, &apods); err != nil {
		log.Fatalf("failed to unmarshal JSON: %v", err)
	}

	repo := repository.NewAPODRepository()

	if err := repo.SaveAPODs(apods); err != nil {
		log.Fatalf("failed to save APODs: %v", err)
	}

	log.Println("APOD data successfully loaded into the database")

}
