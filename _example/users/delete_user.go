package main

import (
	"log"
	"os"

	"github.com/zoom-lib-golang/zoom-lib-golang"
)

func main() {
	var (
		apiKey    = os.Getenv("ZOOM_API_KEY")
		apiSecret = os.Getenv("ZOOM_API_SECRET")
		userID    = os.Getenv("USER_ID")
	)

	zoom.APIKey = apiKey
	zoom.APISecret = apiSecret

	// Delete User by UserID or Email
	opt := zoom.DeleteUserOptions{
		EmailOrID: userID,
	}

	if err := zoom.DeleteUser(opt); err != nil {
		log.Printf("Error: %+v\n", err)
	}
}
