package zoom

import "fmt"

// DeleteUserSSOTokenPath - v2 path for deleting a user's SSO token
const DeleteUserSSOTokenPath = "/users/%s/token"

// DeleteUserSSOTokenOptions are the options to delete a user's SSO token with
type DeleteUserSSOTokenOptions struct {
	EmailOrID string `url:"-"`
}

// DeleteUserSSOToken calls DELETE /users/{userID}/token
func DeleteUserSSOToken(opts DeleteUserSSOTokenOptions) error {
	return defaultClient.DeleteUserSSOToken(opts)
}

// DeleteUserSSOToken calls DELETE /users/{userID}/token
// https://marketplace.zoom.us/docs/api-reference/zoom-api/users/userssotokendelete
func (c *Client) DeleteUserSSOToken(opts DeleteUserSSOTokenOptions) error {
	return c.requestV2(requestV2Opts{
		Method:        Delete,
		Path:          fmt.Sprintf(DeleteUserSSOTokenPath, opts.EmailOrID),
		URLParameters: &opts,
		HeadResponse:  true,
	})
}
