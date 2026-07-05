package service

import (
	"net/url"
	"regexp"
)

var aliasRegex = regexp.MustCompile(`^[a-zA-Z0-9_-]{3,50}$`)

func validateURL(rawURL string) error {
	parsed, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return ErrInvalidURL
	}

	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return ErrInvalidURL
	}

	return nil
}

func validateAlias(alias string) error {
	if alias == "" {
		return nil
	}

	if !aliasRegex.MatchString(alias) {
		return ErrInvalidAlias
	}

	return nil
}
