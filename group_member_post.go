package zoom

import "fmt"

// AddMenbersPath - v2 path for add group members
const AddMenbersPath = "/groups/%s/members"

// AddMemberOptions are details about add group members
type AddMemberOptions struct {
	GroupID string   `json:"-"`
	Members []Member `json:"members"`
}

// ResopnseAddGroupMembers represents response for added member to group
type ResopnseAddGroupMembers struct {
	// IDs has comma-delimited, like 'xxxxxxxxxx,xxxxxxxxxx'
	IDs     string `json:"ids"`
	AddedAt string `json:"added_at"`
}

// AddMembers calls POST /groups/{groupId}/members
func AddMembers(opts AddMemberOptions) (ResopnseAddGroupMembers, error) {
	return defaultClient.AddMembers(opts)
}

// AddMembers calls POST /groups/{groupId}/members
// https://marketplace.zoom.us/docs/api-reference/zoom-api/groups/groupmemberscreate
func (c *Client) AddMembers(opts AddMemberOptions) (ResopnseAddGroupMembers, error) {
	var ret = ResopnseAddGroupMembers{}
	return ret, c.requestV2(requestV2Opts{
		Method:         Post,
		Path:           fmt.Sprintf(AddMenbersPath, opts.GroupID),
		DataParameters: &opts,
		Ret:            &ret,
	})
}
