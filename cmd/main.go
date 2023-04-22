package main

import (
	"fmt"
	"os"

	"github.com/kr/pretty"

	"github.com/hayk99/marvelapp/pkg/domain/service"
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

	comicStorage := store.NewStorage()
	marvelClient := marvel.NewClient(baseURL, publicKey, privateKey)

	comicService := service.NewService(comicStorage, *marvelClient)
	err = comicService.RetrieveNextWeekComics()
	if err != nil {
		fmt.Errorf("error retrieving the comics: %w", err)
		return
	}

	savedComics, err := comicService.ListAllComics()
	if err != nil {
		fmt.Errorf("error listing the comics: %w", err)
		return
	}

	for _, comic := range savedComics {
		pretty.Print(comic)
		fmt.Printf("\n")
	}

	return
}
