package zoom

import "fmt"

// GetRolePath - v2 path for getting a specific role
const GetRolePath = "/roles/%s"

// GetRoleOpts contains options for GetRole
type GetRoleOpts struct {
	ID string `url:"-"`
}

// GetRole calls /roles/{roleId}, searching for a role by ID or email, using the default client
func GetRole(opts GetRoleOpts) (Role, error) {
	return defaultClient.GetRole(opts)
}

// GetRole calls /roles/{roleId}, searching for a role by ID or email, using a specific client
func (c *Client) GetRole(opts GetRoleOpts) (Role, error) {
	var ret = Role{}
	return ret, c.requestV2(requestV2Opts{
		Method:        Get,
		Path:          fmt.Sprintf(GetRolePath, opts.ID),
		URLParameters: opts,
		Ret:           &ret,
	})
}
