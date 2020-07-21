package zoom

import "fmt"

// UpdateUserStatusPath - v2 path for creating a user
const UpdateUserStatusPath = "/users/%s/status"

// UpdateUserStatusURLOptions are the options to create a user with
type UpdateUserStatusURLOptions struct {
	EmailOrID string `url:"-"`
}

// UpdateUserStatusDataOptions are details about a user to create
type UpdateUserStatusDataOptions struct {
	Action UpdateUserStatusAction `json:"action"`
}

// UpdateUserStatus calls POST /users/{userId}/meetings
func UpdateUserStatus(urlOpts UpdateUserStatusURLOptions, dataOpts UpdateUserStatusDataOptions) error {
	return defaultClient.UpdateUserStatus(urlOpts, dataOpts)
}

// UpdateUserStatus calls POST /users
// https://marketplace.zoom.us/docs/api-reference/zoom-api/users/usercreate
func (c *Client) UpdateUserStatus(urlOpts UpdateUserStatusURLOptions, dataOpts UpdateUserStatusDataOptions) error {
	return c.requestV2(requestV2Opts{
		Method:         Put,
		Path:           fmt.Sprintf(UpdateUserStatusPath, urlOpts.EmailOrID),
		DataParameters: &dataOpts,
		HeadResponse:   true,
	})
}
