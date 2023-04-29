package models

import "time"

type ApiCall struct {
	ID            uint `gorm:"primaryKey"`
	ApiCallsCount int
	StartTime     time.Time
	EndTime       time.Time
}
