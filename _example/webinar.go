package main

import (
	"log"
	"os"

	"github.com/himalayan-institute/zoom-lib-golang"
)

// ExampleWebinar contains examples for the /webinar endpoints
func main() {
	var (
		apiKey          = os.Getenv("ZOOM_API_KEY")
		apiSecret       = os.Getenv("ZOOM_API_SECRET")
		email           = os.Getenv("ZOOM_EXAMPLE_EMAIL")
		registrantEmail = os.Getenv("ZOOM_EXAMPLE_REGISTRANT_EMAIL")
	)

	zoom.APIKey = apiKey
	zoom.APISecret = apiSecret
	zoom.Debug = true

	user, err := zoom.GetUser(zoom.GetUserOpts{EmailOrID: email})
	if err != nil {
		log.Fatalf("got error listing users: %+v\n", err)
	}

	fifty := int(50)
	webinars, err := zoom.ListWebinars(zoom.ListWebinarsOptions{
		HostID:   user.ID,
		PageSize: &fifty,
	})
	if err != nil {
		log.Fatalf("got error listing webinars: %+v\n", err)
	}

	log.Printf("Got open webinars: %+v\n", webinars)

	webinars, err = zoom.ListWebinars(zoom.ListWebinarsOptions{
		HostID:   user.ID,
		PageSize: &fifty,
	})
	if err != nil {
		log.Fatalf("got error listing webinars: %+v\n", err)
	}

	log.Printf("Got registration webinars: %+v\n", webinars)

	webinar, err := zoom.GetWebinarInfo(webinars.Webinars[0].ID)

	if err != nil {
		log.Fatalf("got error getting single webinar: %+v\n", err)
	}

	log.Printf("Got single webinars: %+v\n", webinar)

	log.Printf("created at: %s\n", webinar.CreatedAt)
	log.Printf("first occurrence start: %s\n", webinar.Occurrences[0].StartTime)

	customQs := []zoom.CustomQuestion{
		{
			Title: "asdf foo bar",
			Value: "example custom question answer",
		},
	}

	registrantInfo := zoom.WebinarRegistrant{
		WebinarID:       webinar.ID,
		Email:           registrantEmail,
		FirstName:       "Foo",
		LastName:        "Bar",
		CustomQuestions: customQs,
	}

	registrant, err := zoom.RegisterForWebinar(registrantInfo)
	if err != nil {
		log.Fatalf("got error registering a user for webinar %d: %+v\n", webinar.ID, err)
	}

	log.Printf("Got registrant: %+v\n", registrant)

	getRegistrationOpts := zoom.ListWebinarRegistrantsOptions{
		WebinarID: webinar.ID,
	}

	registrationInfo, err := zoom.ListWebinarRegistrants(getRegistrationOpts)
	if err != nil {
		log.Fatalf("got error getting registration info for webinar %d: %+v\n", webinar.ID, err)
	}

	log.Printf("Got registration information: %+v\n", registrationInfo)

	panelists, err := zoom.GetWebinarPanelists(webinar.ID)
	if err != nil {
		log.Fatalf("got error listing webinar panelists for webinar %d: %+v\n", webinar.ID, err)
	}

	log.Printf("Got webinar panelists: %+v\n", panelists)
}
