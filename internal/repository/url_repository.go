package repository

import (
	"database/sql"

	"github.com/shantam-sharma/url-shortner/internal/models"
)

type URLRepository struct {
	db *sql.DB
}

func NewURLRepository(db *sql.DB) *URLRepository {
	return &URLRepository{
		db: db,
	}
}

func (r *URLRepository) Create(url *models.URL) error {
	_, err := r.db.Exec(
		`INSERT INTO urls (original_url, short_code)
		VALUES ($1, $2)`,
		url.OriginalURL,
		url.ShortCode,
	)

	if err != nil {
		return err
	}

	return nil
}
