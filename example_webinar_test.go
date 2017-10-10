package zoom_test

import (
	"log"
	"os"

	"github.com/himalayan-institute/zoom-lib-golang"
)

// ExampleWebinar contains examples for the /webinar endpoints
func ExampleWebinar() {
	apiKey := os.Getenv("ZOOM_API_KEY")
	apiSecret := os.Getenv("ZOOM_API_SECRET")
	email := os.Getenv("ZOOM_EXAMPLE_EMAIL")

	zoom.APIKey = apiKey
	zoom.APISecret = apiSecret
	zoom.Debug = true

	user, err := zoom.GetUserByEmail(zoom.GetUserByEmailOptions{
		Email: email,
	})
	if err != nil {
		log.Fatalf("got error listing users: %+v\n", err)
	}

	log.Printf("Got user: %+v\n", user)
}
