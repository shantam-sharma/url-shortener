package service

import "errors"

var (
	ErrAliasAlreadyExists = errors.New("alias already exists")
	ErrInvalidURL         = errors.New("invalid URL")
	ErrInvalidAlias       = errors.New("invalid alias")
)
