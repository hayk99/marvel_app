package comic

import (
	"github.com/hayk99/marvelapp/pkg/infrastructure/marvel"
)

func ToDomain(comic marvel.ComicDTO) Comic {
	return Comic{
		title:        comic.Title,
		thumbnailURL: comic.Thumbnail.Path + "." + comic.Thumbnail.Extension,
		price:        comic.Prices[0].Value,
	}
}
