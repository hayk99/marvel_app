package store

import (
	"github.com/google/uuid"

	"github.com/hayk99/marvelapp/pkg/domain/comic"
)

func (s *Storage) SaveComic(comic comic.Comic) error {
	s.comics[ID(uuid.New())] = comic
	return nil
}
