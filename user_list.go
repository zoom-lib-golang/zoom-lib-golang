package zoom

// ListUsersPath - v2 path for listing users
const ListUsersPath = "/users"

// ListUsersResponse contains the response from a call to ListUsers
type ListUsersResponse struct {
	TotalRecords int    `json:"total_records"`
	PageCount    int    `json:"page_count"`
	PageNumber   int    `json:"page_number"`
	PageSize     int    `json:"page_size"`
	Users        []User `json:"users"`
}

// ListUsersOptions contains options for ListUsers
type ListUsersOptions struct {
	PageSize   int         `url:"page_size"`
	PageNumber int         `url:"page_number"`
	Status     *UserStatus `url:"status,omitempty"`
}

// ListUsers calls /user/list, listing all users, using the default client
func ListUsers(opts ListUsersOptions) (ListUsersResponse, error) {
	return defaultClient.ListUsers(opts)
}

// ListUsers calls /user/list, listing all users, using client c
func (c *Client) ListUsers(opts ListUsersOptions) (ListUsersResponse, error) {
	var ret = ListUsersResponse{}
	return ret, c.requestV2(requestV2Opts{
		Method:        Get,
		Path:          ListUsersPath,
		URLParameters: opts,
		Ret:           &ret,
	})
}
