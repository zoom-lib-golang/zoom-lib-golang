package zoom

// CreateUserPath - v2 path for creating a user
const CreateUserPath = "/users"

// CreateUserInfo are details about a user to create
type CreateUserInfo struct {
	Email     string   `json:"email"`
	Type      UserType `json:"type"`
	FirstName string   `json:"first_name,omitempty"`
	LastName  string   `json:"last_name,omitempty"`
	Password  string   `json:"password,omitempty"`
}

// CreateUserOptions are the options to create a user with
type CreateUserOptions struct {
	Action   CreateUserAction `json:"action"`
	UserInfo CreateUserInfo   `json:"user_info"`
}

// CreateUser calls POST /users/{userId}/meetings
func CreateUser(opts CreateUserOptions) (User, error) {
	return defaultClient.CreateUser(opts)
}

// CreateUser calls POST /users
// https://marketplace.zoom.us/docs/api-reference/zoom-api/users/usercreate
func (c *Client) CreateUser(opts CreateUserOptions) (User, error) {
	var ret = User{}
	return ret, c.requestV2(requestV2Opts{
		Method:         Post,
		Path:           CreateUserPath,
		DataParameters: &opts,
		Ret:            &ret,
	})
}
