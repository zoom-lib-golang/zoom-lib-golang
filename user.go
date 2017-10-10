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

// ListUsersOptions contains options for ListUsers
type ListUsersOptions struct {
	PageSize   int `url:"page_size"`
	PageNumber int `url:"page_number"`
}

// ListUsers calls /user/list, listing all users, using the default client
func ListUsers(opts ...ListUsersOptions) (ListUsersResponse, error) {
	return defaultClient.ListUsers(opts...)
}

// ListUsers calls /user/list, listing all users, using client c
func (c *Client) ListUsers(opts ...ListUsersOptions) (ListUsersResponse, error) {
	var ret = ListUsersResponse{}
	return ret, request(c, listUsersPath, opts, &ret)
}

// GetUserByEmailOptions contains options for GetUserByEmail
type GetUserByEmailOptions struct {
	Email     string         `url:"email"`
	LoginType *UserLoginType `url:"login_type,omitempty"` // use pointer so it can be null
}

// GetUserByEmail calls /user/getbyemail, searching for a user by email
// address, using the default client
func GetUserByEmail(opts ...GetUserByEmailOptions) (User, error) {
	return defaultClient.GetUserByEmail(opts...)
}

// GetUserByEmail calls /user/getbyemail, searching for a user by email
// address, using client c
func (c *Client) GetUserByEmail(opts ...GetUserByEmailOptions) (User, error) {
	var ret = User{}
	return ret, request(c, getUserByEmailPath, opts, &ret)
}

// GetUser calls /user/get, searching for a user by ID, using the default client
func GetUser(id string) (User, error) {
	return defaultClient.GetUser(id)
}

// GetUser calls /user/get, searching for a user by ID, using client c
func (c *Client) GetUser(id string) (User, error) {
	var ret = User{}
	return ret, request(c, getUserPath, struct {
		ID string `url:"id"`
	}{id}, &ret)
}
