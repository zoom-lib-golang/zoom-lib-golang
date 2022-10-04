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
	)

	zoom.APIKey = apiKey
	zoom.APISecret = apiSecret

	opt := zoom.ListUsersOptions{
		PageSize: 15,
		PageNumber: 1,
	}

	res, err := zoom.ListUsers(opt)

	if err != nil {
		log.Printf("Error: %+v\n", err)
	}
	log.Printf("Resopnse: %+v\n", res)
}
