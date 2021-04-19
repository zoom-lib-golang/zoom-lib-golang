package zoom

import "fmt"

// GetUserPermissionsPath - v2 path for getting a specific user
const GetUserPermissionsPath = "/users/%s/permissions"

// GetUserPermissionsOpts contains options for GetUserPermissions
type GetUserPermissionsOpts struct {
	UserID string `url:"-"`
}

// GetUserPermissions calls /users/{userId}/permissions
func GetUserPermissions(opts GetUserPermissionsOpts) (UserPermissions, error) {
	return defaultClient.GetUserPermissions(opts)
}

// GetUserPermissions calls /users/{userId}/permissions
func (c *Client) GetUserPermissions(opts GetUserPermissionsOpts) (UserPermissions, error) {
	var ret = UserPermissions{}
	return ret, c.requestV2(requestV2Opts{
		Method:        Get,
		Path:          fmt.Sprintf(GetUserPermissionsPath, opts.UserID),
		URLParameters: opts,
		Ret:           &ret,
	})
}

type UserPermissions struct {
	Permissions []string `json:"permissions"`
}
