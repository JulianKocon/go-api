package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title       string
	ReleaseYear int
	Rating      float32
}
