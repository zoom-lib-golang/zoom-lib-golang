package zoom

import "fmt"

// GetGroupPath - v2 path for getting a specific group
const GetGroupPath = "/groups/%s"

// GetGroupOpts contains options for GetGroup
type GetGroupOpts struct {
	ID string `url:"-"`
}

// GetGroup calls /groups/{groupId}, searching for a group by ID or email, using the default client
func GetGroup(opts GetGroupOpts) (Group, error) {
	return defaultClient.GetGroup(opts)
}

// GetGroup calls /groups/{groupId}, searching for a group by ID or email, using a specific client
func (c *Client) GetGroup(opts GetGroupOpts) (Group, error) {
	var ret = Group{}
	return ret, c.requestV2(requestV2Opts{
		Method:        Get,
		Path:          fmt.Sprintf(GetGroupPath, opts.ID),
		URLParameters: opts,
		Ret:           &ret,
	})
}
