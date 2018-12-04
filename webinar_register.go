package zoom

import "fmt"

const (
	// RegisterForWebinarPath - v2 path for registering a user for a webinar
	RegisterForWebinarPath = "/webinars/%d/registrants"

	// ListRegistrantsPath - v2 path for listing panelists for a webinar
	ListRegistrantsPath = "/webinars/%d/registrants"
)

// CustomQuestion is the type for custom questions on registration form
type CustomQuestion struct {
	Title string `json:"title"`
	Value string `json:"value"`
}

// WebinarRegistrant contains options for webinar registration for both
// creating a registration and looking one up. Note that any custom fields are
// strings, and the Title is the actual title entered in Zoom
type WebinarRegistrant struct {
	WebinarID             int                              `json:"-" url:"-"`
	ID                    string                           `json:"id,omitempty"` // for webinar registrant response
	Email                 string                           `json:"email" url:"-"`
	FirstName             string                           `json:"first_name" url:"-"`
	LastName              string                           `json:"last_name" url:"-"`
	Address               string                           `json:"address" url:"-"`
	City                  string                           `json:"city" url:"-"`
	Country               string                           `json:"country" url:"-"` // 2 digit code per https://zoom.github.io/api/#countries
	Zip                   string                           `json:"zip" url:"-"`
	State                 string                           `json:"state" url:"-"`
	Phone                 string                           `json:"phone" url:"-"`
	Industry              string                           `json:"industry" url:"-"`
	Organization          string                           `json:"org" url:"-"`
	JobTitle              string                           `json:"job_title" url:"-"`
	PurchasingTimeFrame   PurchasingTimeFrameType          `json:"purchasing_time_frame" url:"-"`
	RoleInPurchaseProcess PurchaseProcessRoleType          `json:"role_in_purchase_process" url:"-"`
	NumberOfEmployees     NumberOfEmployeesType            `json:"no_of_employees" url:"-"`
	CommentsQuestions     string                           `json:"comments" url:"-"`
	CustomQuestions       []CustomQuestion                 `json:"custom_questions,omitempty" url:"-"` // JSON-encoded []CustomQuestion
	Status                ListWebinarRegistrantsStatusType `json:"status,omitempty" url:"-"`
	CreateTime            *Time                            `json:"create_time" url:"-"`
	JoinURL               *URL                             `json:"join_url" url:"-"`

	/*
		Comma-delimited list of ids. This applies if the webinar is recurring
		with fixed time. The behavior differs based on registration settings:
		1. Register once and attend any - this can be left blank, user is
		   registered for all sessions
		2. Register for each - **behavior unclear**
		3. Register once and attend one or more - **behavior unclear**

		See https://support.zoom.us/hc/en-us/community/posts/115019165043-Behavior-of-occurrence-ids-in-webinar-register-?page=1#community_comment_115004843466
		for more details.
	*/
	OccurrenceIDs string `url:"occurrence_ids,omitempty"`
}

// RegisterForWebinarResponse is the response object returned when registering
// for a webinar
type RegisterForWebinarResponse struct {
	RegistrantID string `json:"registrant_id"`
	WebinarID    int    `json:"id"`
	Topic        string `json:"topic"`
	StartTime    *Time  `json:"start_time"`
	JoinURL      *URL   `json:"join_url"`
}

// RegisterForWebinar registers a user for a webinar, using the default client
func RegisterForWebinar(opts WebinarRegistrant) (RegisterForWebinarResponse, error) {
	return defaultClient.RegisterForWebinar(opts)
}

// RegisterForWebinar registers a user for a webinar, using client c
func (c *Client) RegisterForWebinar(opts WebinarRegistrant) (RegisterForWebinarResponse, error) {
	var ret = RegisterForWebinarResponse{}
	return ret, c.requestV2(requestV2Opts{
		Method:         Post,
		Path:           fmt.Sprintf(RegisterForWebinarPath, opts.WebinarID),
		URLParameters:  opts,
		DataParameters: opts,
		Ret:            &ret,
	})
}

// ListWebinarRegistrantsOptions - options for listing webinar registrants
type ListWebinarRegistrantsOptions struct {
	WebinarID    int                               `url:"-"`
	Status       *ListWebinarRegistrantsStatusType `url:"status,omitempty"`
	PageSize     *int                              `url:"page_size,omitempty"`
	PageNumber   *int                              `url:"page_number,omitempty"`
	OccurrenceID string                            `url:"occurrence_id,omitempty"`
}

// ListWebinarRegistrantsResponse - response for listing webinar registrants
type ListWebinarRegistrantsResponse struct {
	PageCount    int                 `json:"page_count"`
	PageNumber   int                 `json:"page_number"`
	PageSize     int                 `json:"page_size"`
	TotalRecords int                 `json:"total_records"`
	Registrants  []WebinarRegistrant `json:"registrants"`
}

// ListWebinarRegistrants lists webinars using the default client.
func ListWebinarRegistrants(opts ListWebinarRegistrantsOptions) (ListWebinarRegistrantsResponse, error) {
	return defaultClient.ListWebinarRegistrants(opts)
}

// ListWebinarRegistrants lists webinars using client c
func (c *Client) ListWebinarRegistrants(opts ListWebinarRegistrantsOptions) (ListWebinarRegistrantsResponse, error) {
	var ret = ListWebinarRegistrantsResponse{}
	return ret, c.requestV2(requestV2Opts{
		Method:        Get,
		Path:          fmt.Sprintf(RegisterForWebinarPath, opts.WebinarID),
		URLParameters: opts,
		Ret:           &ret,
	})
}
