package zoom

// ListGroupsPath - v2 path for listing groups
const ListGroupsPath = "/groups"

// ListGroupsResponse contains the response from a call to ListGroups
type ListGroupsResponse struct {
	TotalRecords int     `json:"total_records"`
	Groups       []Group `json:"groups"`
}

// ListGroups calls /group/list, listing all groups, using client c
func (c *Client) ListGroups() (ListGroupsResponse, error) {
	var ret = ListGroupsResponse{}
	return ret, c.requestV2(requestV2Opts{
		Method: Get,
		Path:   ListGroupsPath,
		Ret:    &ret,
	})
}
