package models

import "time"

type URL struct {
	ID          int64
	OriginalURL string
	ShortCode   string
	ClickCount  int
	CreatedAt   time.Time
}
