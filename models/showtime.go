package models

import (
	"time"
)

type ShowTime struct {
	ID        int64
	MovieID   int64
	DateTime  time.Time
	Price     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
