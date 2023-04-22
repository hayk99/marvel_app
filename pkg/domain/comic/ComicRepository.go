package comic

type Comic struct {
	title        string
	thumbnailURL string
	price        float32
}

func NewComic(title, thumbnailURL string, price float32) Comic {
	return Comic{
		title:        title,
		thumbnailURL: thumbnailURL,
		price:        price,
	}
}

//go:generate mockgen -source=$GOFILE -destination=./mocks/${GOFILE} -package=mocks
type Repository interface {
	SaveComic(comic Comic) error
	ListAllComics() ([]Comic, error)
}
