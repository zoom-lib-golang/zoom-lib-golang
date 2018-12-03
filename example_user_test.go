package zoom_test

import (
	"log"
	"os"

	"github.com/himalayan-institute/zoom-lib-golang"
)

// ExampleUser contains examples for the /user endpoints
func ExampleUser() {
	apiKey := os.Getenv("ZOOM_API_KEY")
	apiSecret := os.Getenv("ZOOM_API_SECRET")
	email := os.Getenv("ZOOM_EXAMPLE_EMAIL")

	zoom.APIKey = apiKey
	zoom.APISecret = apiSecret
	zoom.Debug = true

	users, err := zoom.ListUsers(zoom.ListUsersOptions{})
	if err != nil {
		log.Fatalf("got error listing users: %+v\n", err)
	}

	log.Printf("Users: %+v\n", users)

	client := zoom.NewClient(apiKey, apiSecret)

	user, err := client.GetUser(zoom.GetUserOpts{EmailOrID: email})
	if err != nil {
		log.Fatalf("got error listing users: %+v\n", err)
	}

	log.Printf("Got user: %+v\n", user)
}
