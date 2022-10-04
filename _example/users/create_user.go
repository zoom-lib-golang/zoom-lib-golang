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
		userEmail = os.Getenv("USER_EMAIL")
	)

	zoom.APIKey = apiKey
	zoom.APISecret = apiSecret

	userInfo := zoom.CreateUserInfo{
		Email: userEmail,
		Type: zoom.Basic,
	}

	opt := zoom.CreateUserOptions {
		Action: zoom.Create,
		UserInfo: userInfo,
	}

	response, err := zoom.CreateUser(opt)

	if err != nil {
		log.Printf("Error: %+v\n", err)
	}
	log.Printf("Resopnse: %+v\n", response)
}
