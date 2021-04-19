package zoom

import "fmt"

// ListGroupMembersPath - v2 path for listing group members
const ListGroupMembersPath = "/groups/%s/members"

// ListGroupMembersResponse contains the response from a call to ListGroupMembers
type ListGroupMembersResponse struct {
	PageCount     int      `json:"page_count"`
	PageNumber    int      `json:"page_number"`
	NextPageToken string   `json:"next_page_token"`
	PageSize      int      `json:"page_size"`
	TotalRecords  int      `json:"total_records"`
	Members       []Member `json:"members"`
}

// ListGroupMembersOptions contains options for ListGroupMembers
type ListGroupMembersOptions struct {
	GroupID       string  `url:"-"`
	PageNumber    int     `json:"page_number"`
	NextPageToken *string `json:"next_page_token,omitempty"`
	PageSize      *int    `json:"page_size,omitempty"`
}

// ListGroupMembers calls /group/list, listing all groups, using client c
func (c *Client) ListGroupMembers(opts ListGroupMembersOptions) (ListGroupMembersResponse, error) {
	var ret = ListGroupMembersResponse{}
	return ret, c.requestV2(requestV2Opts{
		Method:        Get,
		Path:          fmt.Sprintf(ListGroupMembersPath, opts.GroupID),
		URLParameters: opts,
		Ret:           &ret,
	})
}
