package models

import "gorm.io/gorm"

type Reviewer struct {
	gorm.Model
	Movie   Movie  `json: "movie" binding: "required" gorm:"foreignkey:MovieID"`
	MovieID uint64 `json: "-"`
	Reviews []Review
	UserID  uint
}
