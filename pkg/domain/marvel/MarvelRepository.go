package marvel

type Comic struct {
	title        string
	thumbnailURL string
	price        float32
}

//go:generate mockgen -source=$GOFILE -destination=./mocks/${GOFILE} -package=mocks
type Repository interface {
	SaveComic(marvelComic MarvelComic) error
	ListAllComics() ([]Comic, error)
}
