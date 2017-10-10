package main

import (
	"log"
	"os"

	"github.com/himalayan-institute/zoom-lib-golang"
)

func main() {
	apiKey := os.Getenv("ZOOM_API_KEY")
	apiSecret := os.Getenv("ZOOM_API_SECRET")
	email := os.Getenv("ZOOM_EXAMPLE_EMAIL")

	zoom.Debug = true

	client := zoom.NewClient(apiKey, apiSecret)

	response, err := client.GetUserByEmail(zoom.GetUserByEmailOptions{
		Email: email,
	})
	if err != nil {
		log.Fatalf("got error listing users: %+v\n", err)
	}

	log.Printf("Got user: %+v\n", response)
}
