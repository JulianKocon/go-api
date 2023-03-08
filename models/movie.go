package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title       string `json: "title" binding: "required, min=1, max=50" gorm:uniqueIndex`
	ReleaseYear int    `json: "releaseYear" binding: "required, gte=1894, lte=3000"`
	Plot        string `json: "plot"`
}
