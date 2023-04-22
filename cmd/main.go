package main

import (
	"fmt"
	"os"

	"github.com/kr/pretty"

	"github.com/hayk99/marvelapp/pkg/domain/comic"
	"github.com/hayk99/marvelapp/pkg/infrastructure/marvel"
	"github.com/hayk99/marvelapp/pkg/infrastructure/store"
)

func getEnvironmentVars() (string, string, string, error) {
	baseURL := os.Getenv("MARVEL_URL")
	if baseURL == "" {
		return "", "", "", fmt.Errorf("cannot get the endpoint base URL")
	}

	publicKey := os.Getenv("MARVEL_PUBLIC_KEY")
	if publicKey == "" {
		return "", "", "", fmt.Errorf("cannot get the public key ")
	}

	privateKey := os.Getenv("MARVEL_PRIVATE_KEY")
	if privateKey == "" {
		return "", "", "", fmt.Errorf("cannot get the private key")
	}

	return baseURL, publicKey, privateKey, nil
}

func main() {
	baseURL, publicKey, privateKey, err := getEnvironmentVars()
	if err != nil {
		fmt.Printf("error getting the environment vars: %w", err)
	}

	marvel_client := marvel.NewClient(baseURL, publicKey, privateKey)
	comicStorage := store.NewStorage()
	marvelComic, err := marvel_client.GetComicForNextWeek()
	if err != nil {
		fmt.Printf("error getting the comic: %w", err)
	}
	for _, marvelComic := range marvelComic {

		err := comicStorage.SaveComic(comic.ToDomain(marvelComic))
		if err != nil {
			fmt.Printf("error saving the comic: %w", err)
		}
	}

	domainComics, err := comicStorage.ListAllComics()
	if err != nil {
		fmt.Printf("error listing the comics: %w", err)
	}
	for _, domainComic := range domainComics {
		pretty.Print(domainComic)
		fmt.Printf("\n")
	}

	return
}
