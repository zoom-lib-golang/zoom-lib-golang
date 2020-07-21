package zoom

import "fmt"

// UpdateUserStatusPath - v2 path for updating a user's status
const UpdateUserStatusPath = "/users/%s/status"

// UpdateUserStatusURLOptions are the options to identify a user
type UpdateUserStatusURLOptions struct {
	EmailOrID string `url:"-"`
}

// UpdateUserStatusDataOptions are details about a user's status to update
type UpdateUserStatusDataOptions struct {
	Action UpdateUserStatusAction `json:"action"`
}

// UpdateUserStatus calls PUT /users/{userId}/status
func UpdateUserStatus(urlOpts UpdateUserStatusURLOptions, dataOpts UpdateUserStatusDataOptions) error {
	return defaultClient.UpdateUserStatus(urlOpts, dataOpts)
}

// UpdateUserStatus calls PUT /users/{userId}/status
// https://marketplace.zoom.us/docs/api-reference/zoom-api/users/userstatus
func (c *Client) UpdateUserStatus(urlOpts UpdateUserStatusURLOptions, dataOpts UpdateUserStatusDataOptions) error {
	return c.requestV2(requestV2Opts{
		Method:         Put,
		Path:           fmt.Sprintf(UpdateUserStatusPath, urlOpts.EmailOrID),
		DataParameters: &dataOpts,
		HeadResponse:   true,
	})
}
