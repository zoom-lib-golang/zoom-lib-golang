package zoom // Use this file for /webinar endpoints

import "fmt"

const (
	// ListWebinarsPath - v2 lists all webinars
	ListWebinarsPath = "/users/%s/webinars"

	// GetWebinarInfoPath - v2 path for retrieving info on a single webinar
	GetWebinarInfoPath = "/webinars/%d"

	// GetWebinarPanelistsPath - v1 path for listing panelists for a webinar
	GetWebinarPanelistsPath = "/webinar/panelists"
)

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

// GetWebinarPanelistsOptions - options for retrieving webinar panelist info
type GetWebinarPanelistsOptions struct {
	WebinarID int    `url:"id"`
	HostID    string `url:"host_id"`
}

// GetWebinarPanelistsResponse - response from call to /webinar/panelists
type GetWebinarPanelistsResponse struct {
	TotalRecords int               `json:"total_records"`
	Panelists    []WebinarPanelist `json:"panelists"`
}

// GetWebinarPanelists calls /webinar/panelists using the default client
func GetWebinarPanelists(opts ...GetWebinarPanelistsOptions) (GetWebinarPanelistsResponse, error) {
	return defaultClient.GetWebinarPanelists(opts...)
}

// GetWebinarPanelists calls /webinar/panelists using client c
func (c *Client) GetWebinarPanelists(opts ...GetWebinarPanelistsOptions) (GetWebinarPanelistsResponse, error) {
	var ret = GetWebinarPanelistsResponse{}
	return ret, request(c, GetWebinarPanelistsPath, opts, &ret)
}
