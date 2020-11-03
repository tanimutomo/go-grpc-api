package data

import (
	"time"
)

var Articles = map[uint64]Article{
	uint64(1): Article{
		ID:        uint64(1),
		Title:     "title1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	uint64(2): Article{
		ID:        uint64(2),
		Title:     "title2",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	uint64(3): Article{
		ID:        uint64(3),
		Title:     "title3",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}

type Article struct {
	ID        uint64
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
