package zoom

// CreateGroupsPath -v2 path for add groups 
const CreateGroupsPath = "/groups"

// CreateGroupOption - options for creating a group
type CreateGroupOption struct {
	Name string `json:"name"`
}

// CreateGroup calls POST /groups/{groupId}/members
func CreateGroup(opts CreateGroupOption) error {
	return defaultClient.CreateGroup(opts)
}

// CreateGroup calls POST /groups/{groupId}/members
// https://marketplace.zoom.us/docs/api-reference/zoom-api/groups/groupmemberscreate
func (c *Client) CreateGroup(opts CreateGroupOption) error {
	return c.requestV2(requestV2Opts{
		Method:         Post,
		Path:           CreateGroupsPath,
		DataParameters: &opts,
		HeadResponse:   false,
	})
}
