package marvel

type Thumbnail struct {
	Path      string `json:"path"`
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
	Prices    []Price   `json:"prices"`
}

type Data struct {
	Results []MarvelComic `json:"results"`
	Total   int           `json:"total"`
}

type Respose struct {
	Code int  `json:"code"`
	Data Data `json:"data"`
}

func ToDomain(comic MarvelComic) Comic {
	return Comic{
		title:        comic.Title,
		thumbnailURL: comic.Thumbnail.Path + "." + comic.Thumbnail.Extension,
		price:        comic.Prices[0].Value,
	}
}
