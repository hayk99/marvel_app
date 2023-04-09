package store

import (
	"github.com/sirupsen/logrus"

	"github.com/hayk99/marvelapp/pkg/domain/marvel"
)

type Storage struct {
	comics map[string]marvel.Comic
	logger *logrus.Logger
}

func NewStorage() *Storage {
	return &Storage{
		comics: make(map[string]marvel.Comic),
		logger: logrus.New(),
	}
}
