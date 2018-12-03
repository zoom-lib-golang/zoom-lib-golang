package zoom_test

import (
	"encoding/json"
	"log"
	"os"

	"github.com/himalayan-institute/zoom-lib-golang"
)

// ExampleWebinar contains examples for the /webinar endpoints
func ExampleWebinar() {
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

	b, err := json.Marshal(customQs)
	if err != nil {
		log.Fatalf("error marshaling custom Qs to JSON: %s\n", err)
	}

	registrantInfo := zoom.RegisterForWebinarOptions{
		ID:              webinar.ID,
		Email:           registrantEmail,
		FirstName:       "Foo",
		LastName:        "Bar",
		CustomQuestions: string(b),
	}

	registrant, err := zoom.RegisterForWebinar(registrantInfo)
	if err != nil {
		log.Fatalf("got error registering a user for webinar %d: %+v\n", webinar.ID, err)
	}

	log.Printf("Got registrant: %+v\n", registrant)

	getRegistrationOpts := zoom.GetWebinarRegistrationInfoOptions{
		WebinarID: webinar.ID,
		HostID:    user.ID,
	}

	registrationInfo, err := zoom.GetWebinarRegistrationInfo(getRegistrationOpts)
	if err != nil {
		log.Fatalf("got error getting registration info for webinar %d: %+v\n", webinar.ID, err)
	}

	log.Printf("Got registration information: %+v\n", registrationInfo)
}
