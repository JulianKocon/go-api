package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title       string `json: "title" binding: "required"`
	ReleaseYear int    `json: "releaseYear" binding: "required"`
	Plot        string `json: "plot"`
}
