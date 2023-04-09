package marvel

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/hayk99/marvelapp/pkg/domain/marvel"
)

type Client struct {
	baseURL    string
	publicKey  string
	privateKey string
}

func (m *Client) buildGetComicsUrl() string {
	timeStamp := time.Now().Minute()
	hash := md5.Sum([]byte(fmt.Sprint(timeStamp) + m.privateKey + m.publicKey))
	return fmt.Sprintf("%s/v1/public/comics?dateDescriptor=nextWeek&ts=%d&hash=%x&apikey=%s", m.baseURL, timeStamp, hash, m.publicKey)

}

func (m *Client) GetComicForNextWeek() ([]marvel.MarvelComic, error) {
	url := m.buildGetComicsUrl()

	response, err := http.Get(url)
	defer response.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("cannot retrieve comics from marvel service")
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d, cannot retrieve comics successfull: %w", response.StatusCode, err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot read the body: %w", err)
	}

	marvelResponse := marvel.Respose{}
	err = json.Unmarshal(body, &marvelResponse)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal the body: %w", err)
	}

	return marvelResponse.Data.Results, nil
}

func NewClient(baseURL, publicKey, privateKey string) *Client {
	return &Client{
		baseURL,
		publicKey,
		privateKey,
	}
}
