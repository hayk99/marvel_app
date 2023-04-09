package store

import (
	"github.com/hayk99/marvelapp/pkg/domain/marvel"
)

func (s Storage) ListAllComics() ([]marvel.Comic, error) {
	var comics []marvel.Comic
	for _, comic := range s.comics {
		comics = append(comics, comic)
	}
	return comics, nil
}
