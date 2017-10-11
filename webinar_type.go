package zoom // types for /webinar endpoints

// WebinarType is one of a fixed number of possible webinar types
type WebinarType int

const (
	// TypeWebinar is the default webinar type
	TypeWebinar WebinarType = 5

	// TypeRecurringWebinar is a recurring webinar
	TypeRecurringWebinar WebinarType = 6

	// TypeRecurringWebinarFixedTime is a recurring webinar with fixed time
	TypeRecurringWebinarFixedTime WebinarType = 9
)

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
