package zoom // Use this file for /user endpoints

import "fmt"

const (
	listUsersPath = "/users"
	getUserPath   = "/users/%s"
)

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
func ListUsers(opts ...ListUsersOptions) (ListUsersResponse, error) {
	return defaultClient.ListUsers(opts...)
}

// ListUsers calls /user/list, listing all users, using client c
func (c *Client) ListUsers(opts ...ListUsersOptions) (ListUsersResponse, error) {
	var ret = ListUsersResponse{}
	return ret, c.requestV2(requestV2Opts{
		Method:        Get,
		Path:          listUsersPath,
		URLParameters: opts,
		Ret:           &ret,
	})
}

// GetUserOpts contains options for GetUser
type GetUserOpts struct {
	EmailOrID string         `url:"userId"`
	LoginType *UserLoginType `url:"login_type,omitempty"` // use pointer so it can be null
}

// GetUser calls /users/{userId}, searching for a user by ID or email, using the default client
func GetUser(opts GetUserOpts) (User, error) {
	return defaultClient.GetUser(opts)
}

// GetUser calls /users/{userId}, searching for a user by ID or email, using a specific client
func (c *Client) GetUser(opts GetUserOpts) (User, error) {
	var ret = User{}
	return ret, c.requestV2(requestV2Opts{
		Method:        Get,
		Path:          fmt.Sprintf(getUserPath, opts.EmailOrID),
		URLParameters: opts,
		Ret:           &ret,
	})
}
