package zoom // Use this file for /webinar endpoints

const (
	listWebinarsPath             = "/webinar/list"
	listRegistrationWebinarsPath = "/webinar/list/registration"
	getWebinarInfoPath           = "/webinar/get"
	registerForWebinarPath       = "/webinar/register"
	getWebinarRegistrationInfo   = "/webinar/registration"
)

// ListWebinarsResponse contains the response from a call to ListWebinars
type ListWebinarsResponse struct {
	PageCount    int       `json:"page_count"`
	TotalRecords int       `json:"total_records"`
	PageNumber   int       `json:"page_number"`
	PageSize     int       `json:"page_size"`
	Webinars     []Webinar `json:"webinars"`
}

// ListWebinarsOptions contains options for ListWebinars
type ListWebinarsOptions struct {
	HostID     string `url:"host_id"`
	PageSize   *int   `url:"page_size,omitempty"`
	PageNumber *int   `url:"page_number,omitempty"`
}

// ListWebinars calls /webinar/list, listing all webinars that don't require
// registration, using the default client
func ListWebinars(opts ...ListWebinarsOptions) (ListWebinarsResponse, error) {
	return defaultClient.ListWebinars(opts...)
}

// ListWebinars calls /webinar/list, listing all webinars that don't require
// registration, using the client c
func (c *Client) ListWebinars(opts ...ListWebinarsOptions) (ListWebinarsResponse, error) {
	var ret = ListWebinarsResponse{}
	return ret, request(c, listWebinarsPath, opts, &ret)
}

// ListRegistrationWebinars calls /webinar/list, listing all webinars that don't require
// registration, using the default client
func ListRegistrationWebinars(opts ...ListWebinarsOptions) (ListWebinarsResponse, error) {
	return defaultClient.ListRegistrationWebinars(opts...)
}

// ListRegistrationWebinars calls /webinar/list, listing all webinars that don't require
// registration, using the client c
func (c *Client) ListRegistrationWebinars(opts ...ListWebinarsOptions) (ListWebinarsResponse, error) {
	var ret = ListWebinarsResponse{}
	return ret, request(c, listRegistrationWebinarsPath, opts, &ret)
}

// GetWebinarInfoOptions contains options for GetWebinarInfo
type GetWebinarInfoOptions struct {
	ID     int    `url:"id"`
	HostID string `url:"host_id"`
}

// GetWebinarInfo calls /webinar/get, listing a single webinar, using the
// default client
func GetWebinarInfo(opts ...GetWebinarInfoOptions) (Webinar, error) {
	return defaultClient.GetWebinarInfo(opts...)
}

// GetWebinarInfo calls /webinar/get, listing a single webinar, using client c
func (c *Client) GetWebinarInfo(opts ...GetWebinarInfoOptions) (Webinar, error) {
	var ret = Webinar{}
	return ret, request(c, getWebinarInfoPath, opts, &ret)
}

// CustomQuestion is the type for custom questions on registration form
type CustomQuestion struct {
	Title string `json:"title"`
	Value string `json:"value"`
}

// RegisterForWebinarOptions contains options for webinar registration. Note
// that any custom fields are strings, and the Title is the actual title
// entered in Zoom
type RegisterForWebinarOptions struct {
	ID                    int                     `url:"id"`
	Email                 string                  `url:"email"`
	FirstName             string                  `url:"first_name"`
	LastName              string                  `url:"last_name"`
	Address               string                  `url:"address"`
	City                  string                  `url:"city"`
	Country               string                  `url:"country"` // 2 digit code per https://zoom.github.io/api/#countries
	Zip                   string                  `url:"zip"`
	State                 string                  `url:"state"`
	Phone                 string                  `url:"phone"`
	Industry              string                  `url:"industry"`
	Organization          string                  `url:"org"`
	JobTitle              string                  `url:"job_title"`
	PurchasingTimeFrame   PurchasingTimeFrameType `url:"purchasing_time_frame"`
	RoleInPurchaseProcess PurchaseProcessRoleType `url:"role_in_purchase_process"`
	NumberOfEmployees     NumberOfEmployeesType   `url:"no_of_employees"`
	CommentsQuestions     string                  `url:"comments"`
	CustomQuestions       string                  `url:"custom_questions,omitempty"` // JSON-encoded []CustomQuestion
	Language              string                  `url:"language"`

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

// RegisterForWebinar calls /webinar/register using the default client
func RegisterForWebinar(opts ...RegisterForWebinarOptions) (RegisterForWebinarResponse, error) {
	return defaultClient.RegisterForWebinar(opts...)
}

// RegisterForWebinar calls /webinar/register using client c
func (c *Client) RegisterForWebinar(opts ...RegisterForWebinarOptions) (RegisterForWebinarResponse, error) {
	var ret = RegisterForWebinarResponse{}
	return ret, request(c, registerForWebinarPath, opts, &ret)
}

// GetWebinarRegistrationInfoOptions is the response object returned when registering
// for a webinar
type GetWebinarRegistrationInfoOptions struct {
	WebinarID    int    `url:"id"`
	HostID       string `url:"host_id"`
	OccurrenceID string `url:"occurrence_id"`
	PageSize     *int   `url:"page_size"`
	PageNumber   *int   `url:"page_number"`
}

// GetWebinarRegistrationInfoResponse is the response data for a call to
// /webinar/registration for getting webinar registrantion information
type GetWebinarRegistrationInfoResponse struct {
	PageCount    int                 `json:"page_count"`
	PageNumber   int                 `json:"page_number"`
	PageSize     int                 `json:"page_size"`
	TotalRecords int                 `json:"total_records"`
	Registrants  []WebinarRegistrant `json:"attendees"`
}

// GetWebinarRegistrationInfo calls /webinar/registration using the default client
func GetWebinarRegistrationInfo(opts ...GetWebinarRegistrationInfoOptions) (GetWebinarRegistrationInfoResponse, error) {
	return defaultClient.GetWebinarRegistrationInfo(opts...)
}

// GetWebinarRegistrationInfo calls /webinar/registration using client c
func (c *Client) GetWebinarRegistrationInfo(opts ...GetWebinarRegistrationInfoOptions) (GetWebinarRegistrationInfoResponse, error) {
	var ret = GetWebinarRegistrationInfoResponse{}
	return ret, request(c, getWebinarRegistrationInfo, opts, &ret)
}
