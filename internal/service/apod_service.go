package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	model "nied-science/internal/model"
	"nied-science/internal/repository"
	"os"
	"time"
)

type APODService interface {
	FetchAndSaveAPOD() error
}

type apodService struct {
	repo repository.APODRepository
}

func NewAPODService(repo repository.APODRepository) APODService {
	return &apodService{repo: repo}
}

func (s *apodService) FetchAndSaveAPOD() error {
	apod, err := fetchAPOD()
	if err != nil {
		return err
	}
	return s.repo.SaveAPOD(apod)
}

func fetchAPOD() (*model.APOD, error) {
	apiKeyNasa := os.Getenv("API_KEY_NASA")
	if apiKeyNasa == "" {
		return nil, fmt.Errorf("API_KEY_NASA não está definida")
	}

	now := time.Now()
	formattedDate := now.Format("2006-01-02")

	apodURL := fmt.Sprintf("https://api.nasa.gov/planetary/apod?api_key=%s&date=%s", apiKeyNasa, formattedDate)

	resp, err := http.Get(apodURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var apod model.APOD
	if err := json.Unmarshal(body, &apod); err != nil {
		return nil, err
	}
	return &apod, nil
}
