package zoom // Use this file for /user endpoints

const (
	listUsersPath      = "/user/list"
	getUserPath        = "/user/get"
	getUserByEmailPath = "/user/getbyemail"
)

// ListUsersResponse contains the response from a call to ListUsers
type ListUsersResponse struct {
	PageCount    int    `json:"page_count"`
	TotalRecords int    `json:"total_records"`
	PageNumber   int    `json:"page_number"`
	PageSize     int    `json:"page_size"`
	Users        []User `json:"users"`
}

// ListUsersOptions contains options to ListUsers
type ListUsersOptions struct {
	PageSize   int `url:"page_size"`
	PageNumber int `url:"page_number"`
}

// ListUsers calls /user/list using the default client
func ListUsers(opts ...ListUsersOptions) (ListUsersResponse, error) {
	return defaultClient.ListUsers(opts...)
}

// ListUsers calls /user/list for an instantiated client
func (c *Client) ListUsers(opts ...ListUsersOptions) (ListUsersResponse, error) {
	var ret = ListUsersResponse{}
	return ret, executeRequest(c, listUsersPath, opts, &ret)
}
