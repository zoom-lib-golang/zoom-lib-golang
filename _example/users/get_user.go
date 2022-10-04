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
		userEmail = os.Getenv("USER_EMAIL")
	)

	zoom.APIKey = apiKey
	zoom.APISecret = apiSecret

	// Get User info by UserID
	opt := zoom.GetUserOpts{
		EmailOrID: userID,
	}

	response, err := zoom.GetUser(opt)

	if err != nil {
		log.Printf("Error: %+v\n", err)
	}
	log.Printf("Resopnse: %+v\n", response)

	// Get User info by Email
	opt = zoom.GetUserOpts{
		EmailOrID: userEmail,
	}

	response, err = zoom.GetUser(opt)

	if err != nil {
		log.Printf("Error: %+v\n", err)
	}
	log.Printf("Resopnse: %+v\n", response)
}
