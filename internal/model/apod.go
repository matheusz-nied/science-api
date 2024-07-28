package model

type APOD struct {
	Date        string `json:"date"`
	Explanation string `json:"explanation"`
	Hdurl       string `json:"hdurl"`
	MediaType   string `json:"media_type"`
	Url         string `json:"url"`
	Title       string `json:"title"`
}
