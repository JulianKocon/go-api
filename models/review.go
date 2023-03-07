package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	Reviewer string  `json: "reviewer" binding: "min=5, max=50`
	Rating   float32 `json: "rating"`
	Review   string  `json: "review"  binding: "max=500`
	Movie    Movie   `json: "movie" binding: "required" gorm:"foreignkey:MovieID"`
	MovieID  uint64  `json: "-"`
}
