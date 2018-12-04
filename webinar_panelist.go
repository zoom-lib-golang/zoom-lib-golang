package zoom

import "fmt"

const (
	// GetWebinarPanelistsPath - v2 path for listing panelists for a webinar
	GetWebinarPanelistsPath = "/webinars/%d/panelists"
)

// WebinarPanelist contains information returned by /webinar/panelists
type WebinarPanelist struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	JoinURL *URL   `json:"join_url"`
}

// GetWebinarPanelistsResponse - response from call to /webinar/panelists
type GetWebinarPanelistsResponse struct {
	TotalRecords int               `json:"total_records"`
	Panelists    []WebinarPanelist `json:"panelists"`
}

// GetWebinarPanelists calls /webinar/panelists using the default client
func GetWebinarPanelists(webinarID int) (GetWebinarPanelistsResponse, error) {
	return defaultClient.GetWebinarPanelists(webinarID)
}

// GetWebinarPanelists calls /webinar/panelists using client c
func (c *Client) GetWebinarPanelists(webinarID int) (GetWebinarPanelistsResponse, error) {
	var ret = GetWebinarPanelistsResponse{}
	return ret, c.requestV2(requestV2Opts{
		Method: Get,
		Path:   fmt.Sprintf(GetWebinarPanelistsPath, webinarID),
		Ret:    &ret,
	})
}
