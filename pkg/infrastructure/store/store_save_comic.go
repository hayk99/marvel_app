package store

import (
	"fmt"

	"github.com/hayk99/marvelapp/pkg/domain/marvel"
)

func (s *Storage) SaveComic(marvelComic marvel.MarvelComic) error {
	s.comics[fmt.Sprint(marvelComic.ID)] = marvel.ToDomain(marvelComic)
	return nil
}
