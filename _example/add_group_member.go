package main

import (
	"log"
	"os"

	"github.com/himalayan-institute/zoom-lib-golang"
)

func main() {
	var (
		apiKey    = os.Getenv("ZOOM_API_KEY")
		apiSecret = os.Getenv("ZOOM_API_SECRET")
		userID    = os.Getenv("USER_ID")
		groupID   = os.Getenv("GROUP_ID")
	)

	zoom.APIKey = apiKey
	zoom.APISecret = apiSecret

	addMemberopts := zoom.AddMemberOptions{
		GroupID: groupID,
		Members: []zoom.Member{
			{
				ID: userID,
			},
		},
	}

	member, err := zoom.AddMembers(addMemberopts)
	if err != nil {
		log.Printf("Got add members: %+v\n", err)
	}
	log.Printf("Resopnse add members: %+v\n", member)
}
