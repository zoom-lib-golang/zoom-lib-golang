package zoom // types for /webinar endpoints

// Webinar represents a webinar object
type Webinar struct {
	UUID                      string              `json:"uuid"`
	ID                        int                 `json:"id"`
	StartURL                  string              `json:"start_url"`
	JoinURL                   string              `json:"join_url"`
	RegistrationURL           string              `json:"registration_url"`
	CreatedAt                 *Time               `json:"created_at"`
	HostID                    string              `json:"host_id"`
	Topic                     string              `json:"topic"`
	Type                      WebinarType         `json:"type"`
	StartTime                 *Time               `json:"start_time"`
	Duration                  int                 `json:"duration"`
	Timezone                  string              `json:"timezone"`
	Agenda                    string              `json:"agenda"`
	OptionStartType           string              `json:"option_start_type"`
	OptionAudio               string              `json:"option_audio"`
	OptionEnforceLogin        bool                `json:"option_enforce_login"`
	OptionEnforceLoginDomains string              `json:"option_enforce_login_domains"`
	OptionAlternativeHosts    string              `json:"option_alternative_hosts"`
	Status                    int                 `json:"status"`
	Occurrences               []WebinarOccurrence `json:"occurrences"`
}

// WebinarOccurrence contains recurrence data for recurring webinars
type WebinarOccurrence struct {
	OccurrenceID string `json:"occurrence_id"`
	StartTime    *Time  `json:"start_time"`
	Duration     int    `json:"duration"`
}

// WebinarRegistrant is returned when retrieving information about webinar
// attendees from calls to /webinar/registration
type WebinarRegistrant struct {
	ID               string `json:"id"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	CreateTime       *Time  `json:"create_time"`
	Approval         string `json:"approval"`
	Email            string `json:"email"`
	RegistrationURL  *URL   `json:"registration_url"`
	RegistrationTime *Time  `json:"registration_time"`
}

// WebinarPanelist contains information returned by /webinar/panelists
type WebinarPanelist struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	JoinURL *URL   `json:"join_url"`
}
