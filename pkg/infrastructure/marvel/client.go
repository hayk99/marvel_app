package marvel

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"time"
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

func (m *Client) GetComicForNextWeek() (string, error) {
	url := m.buildGetComicsUrl()

	response, err := http.Get(url)
	defer response.Body.Close()
	if err != nil {
		return "", fmt.Errorf("cannot retrieve comics from marvel service")
	}

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status code: %d, cannot retrieve comics successfull: %w", response.StatusCode, err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("cannot read the body: %w", err)
	}

	return bytes.NewBuffer(body).String(), nil
}

func NewClient(baseURL, publicKey, privateKey string) *Client {
	return &Client{
		baseURL,
		publicKey,
		privateKey,
	}
}
