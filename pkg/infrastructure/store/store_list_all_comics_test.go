package store_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/hayk99/marvelapp/pkg/domain/marvel"
	"github.com/hayk99/marvelapp/pkg/infrastructure/store"
)

var _ = Describe("StoreListAllComics", func() {

	var (
		comicStore  *store.Storage
		marvelComic []marvel.MarvelComic
	)
	BeforeEach(func() {
		comicStore = store.NewStorage()
		marvelComic = []marvel.MarvelComic{
			{
				ID:    1,
				Title: "Comic 1",
				Thumbnail: marvel.Thumbnail{
					Path:      "www.somePath.com/to/image",
					Extension: "jpg",
				},
				Prices: []marvel.Price{
					{
						Type:  "printPrice",
						Value: 1.99,
					},
				},
			},
			{
				ID:    2,
				Title: "Comic 2",
				Thumbnail: marvel.Thumbnail{
					Path:      "www.somePath.com/to/image2",
					Extension: "jpg",
				},
				Prices: []marvel.Price{
					{
						Type:  "printPrice",
						Value: 1.99,
					},
				},
			},
		}
	})

	Context("Given a comic", func() {
		It("saves the comic", func() {
			err := comicStore.SaveComic(marvelComic[0])
			Expect(err).ToNot(HaveOccurred())
			err = comicStore.SaveComic(marvelComic[1])
			Expect(err).ToNot(HaveOccurred())

			comics, err := comicStore.ListAllComics()
			Expect(err).ToNot(HaveOccurred())
			Expect(comics).To(HaveLen(2))
		})
	})

})
