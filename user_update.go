package zoom

import "fmt"

// UpdateUserPath - v2 path for update a user
const UpdateUserPath = "/users/%s"

// UpdateUserOpts contains options for UpdateUser
type UpdateUserOpts struct {
	EmailOrID            string         `url:"-"`
	LoginType            *UserLoginType `url:"login_type,omitempty"` // use pointer so it can be null
	RemoveTspCredentials bool           `url:"remove_tsp_credentials,omitempty"`
}

// UpdateUser calls /users/{userId}, update for a user by ID or email, using the default client
func UpdateUser(user User, opts UpdateUserOpts) error {
	return defaultClient.UpdateUser(user, opts)
}

// UpdateUser calls /users/{userId}, update for a user by ID or email, using the specific client
func (c *Client) UpdateUser(user User, opts UpdateUserOpts) error {
	return c.requestV2(requestV2Opts{
		Method:         Patch,
		Path:           fmt.Sprintf(UpdateUserPath, opts.EmailOrID),
		URLParameters:  opts,
		DataParameters: user,
		HeadResponse:   true,
	})
}
