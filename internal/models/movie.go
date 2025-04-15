package models

type Movie struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	Genre       string `json:"genre"`
	ReleaseDate string `json:"release_date"`
}
