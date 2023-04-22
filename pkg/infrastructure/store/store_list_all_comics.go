package store

import (
	"github.com/hayk99/marvelapp/pkg/domain/comic"
)

func (s Storage) ListAllComics() ([]comic.Comic, error) {
	var comics []comic.Comic
	for _, comic := range s.comics {
		comics = append(comics, comic)
	}
	return comics, nil
}
