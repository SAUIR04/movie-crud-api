package models

type Movie struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	Genre       string `json:"genre"`
	ReleaseDate string `json:"release_date"`
}
