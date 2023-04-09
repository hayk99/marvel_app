package main

import (
	"fmt"
	"os"

	"github.com/hayk99/marvelapp/pkg/infrastructure/marvel"
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

}
