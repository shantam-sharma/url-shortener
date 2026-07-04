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

func (r *URLRepository) GetByShortCode(code string) (*models.URL, error) {
	url := &models.URL{}

	query := `
		SELECT id, original_url, short_code, click_count, created_at
		FROM urls
		WHERE short_code = $1
	`

	err := r.db.QueryRow(query, code).Scan(
		&url.ID,
		&url.OriginalURL,
		&url.ShortCode,
		&url.ClickCount,
		&url.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return url, nil
}

func (r *URLRepository) IncrementClickCount(code string) error {
	query := `
		UPDATE urls
		SET click_count = click_count + 1
		WHERE short_code = $1
	`

	_, err := r.db.Exec(query, code)
	if err != nil {
		return err
	}

	return nil
}

func (r *URLRepository) AliasExists(alias string) (bool, error) {
	var exists bool

	query := `
		SELECT EXISTS(
			SELECT 1
			FROM urls
			WHERE short_code = $1
		)
	`

	err := r.db.QueryRow(query, alias).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
