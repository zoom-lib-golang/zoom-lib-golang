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
