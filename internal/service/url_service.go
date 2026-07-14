package service

import (
	"crypto/rand"
	"math/big"

	"github.com/shantam-sharma/url-shortener/internal/models"
	"github.com/shantam-sharma/url-shortener/internal/repository"
)

type URLService struct {
	repo *repository.URLRepository
}

func NewURLService(repo *repository.URLRepository) *URLService {
	return &URLService{
		repo: repo,
	}
}

func generateShortCode(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	shortCode := make([]byte, length)

	for i := range shortCode {
		randomIndex, err := rand.Int(
			rand.Reader,
			big.NewInt(int64(len(charset))),
		)

		if err != nil {
			panic(err)
		}

		shortCode[i] = charset[randomIndex.Int64()]
	}

	return string(shortCode)
}

func (s *URLService) Create(originalURL, alias string) (*models.URL, error) {

	if err := validateURL(originalURL); err != nil {
		return nil, err
	}

	if err := validateAlias(alias); err != nil {
		return nil, err
	}

	var shortCode string

	if alias != "" {
		exists, err := s.repo.AliasExists(alias)
		if err != nil {
			return nil, err
		}

		if exists {
			return nil, ErrAliasAlreadyExists
		}

		shortCode = alias
	} else {
		shortCode = generateShortCode(6)
	}

	url := &models.URL{
		OriginalURL: originalURL,
		ShortCode:   shortCode,
	}

	if err := s.repo.Create(url); err != nil {
		return nil, err
	}

	return url, nil
}

func (s *URLService) Resolve(shortCode string) (*models.URL, error) {
	url, err := s.repo.GetByShortCode(shortCode)
	if err != nil {
		return nil, err
	}

	err = s.repo.IncrementClickCount(shortCode)
	if err != nil {
		return nil, err
	}

	return url, nil
}
