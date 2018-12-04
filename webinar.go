package zoom // Use this file for /webinar endpoints

import "fmt"

const (
	// ListWebinarsPath - v2 lists all webinars
	ListWebinarsPath = "/users/%s/webinars"

	// GetWebinarInfoPath - v2 path for retrieving info on a single webinar
	GetWebinarInfoPath = "/webinars/%d"
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

// ListWebinarsResponse contains the response from a call to ListWebinars
type ListWebinarsResponse struct {
	PageCount    int       `json:"page_count"`
	TotalRecords int       `json:"total_records"`
	PageNumber   int       `json:"page_number"`
	PageSize     int       `json:"page_size"`
	Webinars     []Webinar `json:"webinars"`
}

// ListWebinarsOptions contains options for ListWebinars. Also accepts email address for HostID
type ListWebinarsOptions struct {
	HostID     string `url:"-"`
	PageSize   *int   `url:"page_size,omitempty"`
	PageNumber *int   `url:"page_number,omitempty"`
}

// ListWebinars calls /webinar/list, listing all webinars that don't require
// registration, using the default client
func ListWebinars(opts ListWebinarsOptions) (ListWebinarsResponse, error) {
	return defaultClient.ListWebinars(opts)
}

// ListWebinars calls /webinar/list, listing all webinars that don't require
// registration, using the client c
func (c *Client) ListWebinars(opts ListWebinarsOptions) (ListWebinarsResponse, error) {
	var ret = ListWebinarsResponse{}
	return ret, c.requestV2(requestV2Opts{
		Method:        Get,
		Path:          fmt.Sprintf(ListWebinarsPath, opts.HostID),
		URLParameters: &opts,
		Ret:           &ret,
	})
}

// GetWebinarInfo gets into about a single webinar, using the default client
func GetWebinarInfo(webinarID int) (Webinar, error) {
	return defaultClient.GetWebinarInfo(webinarID)
}

// GetWebinarInfo gets into about a single webinar, using client c
func (c *Client) GetWebinarInfo(webinarID int) (Webinar, error) {
	var ret = Webinar{}
	return ret, c.requestV2(requestV2Opts{
		Method: Get,
		Path:   fmt.Sprintf(GetWebinarInfoPath, webinarID),
		Ret:    &ret,
	})
}
