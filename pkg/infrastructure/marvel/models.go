package marvel

type Thumbnail struct {
	Path      string `json:"path"`
	Extension string `json:"extension"`
}

type Price struct {
	Type  string  `json:"type"`
	Value float32 `json:"price"`
}

type ComicDTO struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Thumbnail Thumbnail `json:"thumbnail"`
	Prices    []Price   `json:"prices"`
}

type Data struct {
	Results []ComicDTO `json:"results"`
	Total   int        `json:"total"`
}

type Respose struct {
	Code int  `json:"code"`
	Data Data `json:"data"`
}
