package marvel

type Thumbnail struct {
	Path      string `json:"url"`
	Extension string `json:"extension"`
}

type Price struct {
	Type  string  `json:"type"`
	Value float32 `json:"price"`
}

type MarvelComic struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Thumbnail Thumbnail `json:"thumbnail"`
	Price     Price     `json:"prices"`
}

func ToDomain(comic MarvelComic) Comic {
	return Comic{
		title:        comic.Title,
		thumbnailURL: comic.Thumbnail.Path + "." + comic.Thumbnail.Extension,
		price:        comic.Price.Value,
	}
}
