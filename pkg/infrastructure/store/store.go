package store

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"github.com/hayk99/marvelapp/pkg/domain/comic"
)

type ID uuid.UUID

type Storage struct {
	comics map[ID]comic.Comic
	logger *logrus.Logger
}

func NewStorage() *Storage {
	return &Storage{
		comics: make(map[ID]comic.Comic),
		logger: logrus.New(),
	}
}
