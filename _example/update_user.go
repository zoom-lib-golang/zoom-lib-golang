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

	// Get User info by UserID
	getOpt := zoom.GetUserOpts{
		EmailOrID: userID,
	}

	user, err := zoom.GetUser(getOpt)

	if err != nil {
		log.Printf("Error: %+v\n\n", err)
	}
	log.Printf("Resopnse: %+v\n\n", user)

	// Update User info by Email
	updateOpt := zoom.UpdateUserOpts{
		EmailOrID: user.Email,
	}

	user.FirstName = "Zoom"

	if err = zoom.UpdateUser(user, updateOpt); err != nil {
		log.Printf("Error: %+v\n\n", err)
	}
}
