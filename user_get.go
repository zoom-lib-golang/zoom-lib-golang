package zoom

import "fmt"

// GetUserPath - v2 path for getting a specific user
const GetUserPath = "/users/%s"

// GetUserOpts contains options for GetUser
type GetUserOpts struct {
	EmailOrID string         `url:"-"`
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
		Path:          fmt.Sprintf(GetUserPath, opts.EmailOrID),
		URLParameters: opts,
		Ret:           &ret,
	})
}
