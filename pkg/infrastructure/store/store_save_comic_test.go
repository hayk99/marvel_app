package store_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/hayk99/marvelapp/pkg/domain/comic"
	"github.com/hayk99/marvelapp/pkg/infrastructure/store"
)

var _ = Describe("StoreSaveComic", func() {

	var (
		comicStore *store.Storage
		newEntry   comic.Comic
	)
	BeforeEach(func() {
		comicStore = store.NewStorage()
		newEntry = comic.NewComic("Comic 1", "www.somePath.com/to/image.jpg", 1.99)
	})

	Context("Given a comic", func() {
		It("saves the comic", func() {
			err := comicStore.SaveComic(newEntry)
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
