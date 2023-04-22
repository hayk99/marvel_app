package marvel_test

import (
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/hayk99/marvelapp/pkg/infrastructure/marvel"
)

const baseURL = "https://gateway.marvel.com:443/"

var _ = Describe("Client", func() {
	var (
		client           *marvel.Client
		marvelPublicKey  = os.Getenv("MARVEL_PUBLIC_KEY")
		marvelPrivateKey = os.Getenv("MARVEL_PRIVATE_KEY")
	)

	BeforeEach(func() {
		client = marvel.NewClient(baseURL, marvelPublicKey, marvelPrivateKey)
	})

	Context("Given the URL", func() {
		It("retrieves the comics for the week", func() {
			comicsList, err := client.GetComicForNextWeek()
			Expect(err).ToNot(HaveOccurred())
			Expect(comicsList).ToNot(BeEmpty())
		})
	})
})
