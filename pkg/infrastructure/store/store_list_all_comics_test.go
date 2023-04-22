package store_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/hayk99/marvelapp/pkg/domain/comic"
	"github.com/hayk99/marvelapp/pkg/infrastructure/store"
)

var _ = Describe("StoreListAllComics", func() {

	var (
		comicStore *store.Storage
		comics     []comic.Comic
	)
	BeforeEach(func() {
		comicStore = store.NewStorage()

		comics[0] = comic.NewComic("Comic 1", "www.somePath.com/to/image.jpg", 1.99)
		comics[1] = comic.NewComic("Comic 2", "www.somePath.com/to/image2.jpg", 1.99)
	})

	Context("Given a comic", func() {
		It("saves the comic", func() {
			err := comicStore.SaveComic(comics[0])
			Expect(err).ToNot(HaveOccurred())
			err = comicStore.SaveComic(comics[1])
			Expect(err).ToNot(HaveOccurred())

			comics, err := comicStore.ListAllComics()
			Expect(err).ToNot(HaveOccurred())
			Expect(comics).To(HaveLen(2))
		})
	})

})
