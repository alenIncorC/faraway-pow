package repository

import "github.com/alenIncorC/faraway-pow/faraway-tcp-server/internal/repository/file"

type Quotes interface {
	GetQuote() (string, error)
}

type Repositories struct {
	Quotes
}

func NewRepositories() Repositories {
	return Repositories{
		Quotes: file.NewQuote(),
	}
}
