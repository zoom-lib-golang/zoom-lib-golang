package zoom

import "fmt"

// ListRoleMembersPath - v2 path for listing role members
const ListRoleMembersPath = "/roles/%s/members"

// ListRoleMembersResponse contains the response from a call to ListRoleMembers
type ListRoleMembersResponse struct {
	PageCount     int      `json:"page_count"`
	PageNumber    int      `json:"page_number"`
	NextPageToken string   `json:"next_page_token"`
	PageSize      int      `json:"page_size"`
	TotalRecords  int      `json:"total_records"`
	Members       []Member `json:"members"`
}

// ListRoleMembersOptions contains options for ListRoleMembers
type ListRoleMembersOptions struct {
	RoleID        string `url:"-"`
	PageNumber    int    `json:"page_number"`
	NextPageToken string `json:"next_page_token"`
	PageSize      *int   `json:"page_size,omitempty"`
}

// Member is a member of a group or role
type Member struct {
	ID         string   `json:"id"`
	Email      string   `json:"email"`
	FirstName  string   `json:"first_name"`
	LastName   string   `json:"last_name"`
	Type       UserType `json:"type"`
	Department string   `json:"department"`
}

// ListRoleMembers calls /role/list, listing all roles, using client c
func (c *Client) ListRoleMembers(opts ListRoleMembersOptions) (ListRoleMembersResponse, error) {
	var ret = ListRoleMembersResponse{}
	return ret, c.requestV2(requestV2Opts{
		Method:        Get,
		Path:          fmt.Sprintf(ListRoleMembersPath, opts.RoleID),
		URLParameters: opts,
		Ret:           &ret,
	})
}
