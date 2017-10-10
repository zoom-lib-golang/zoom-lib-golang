package main

import (
	"log"
	"os"

	"github.com/himalayan-institute/zoom-lib-golang"
)

func main() {
	zoom.APIKey = os.Getenv("ZOOM_API_KEY")
	zoom.APISecret = os.Getenv("ZOOM_API_SECRET")
	zoom.Debug = true

	response, err := zoom.ListUsers()
	if err != nil {
		log.Fatalf("got error listing users: %+v\n", err)
	}

	log.Printf("Users: %+v\n", response)
}
