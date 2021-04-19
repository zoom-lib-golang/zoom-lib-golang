package zoom

// ListRolesPath - v2 path for listing roles
const ListRolesPath = "/roles"

// ListRolesResponse contains the response from a call to ListRoles
type ListRolesResponse struct {
	TotalRecords int    `json:"total_records"`
	Roles        []Role `json:"roles"`
}

// ListRoles calls /role/list, listing all roles, using client c
func (c *Client) ListRoles() (ListRolesResponse, error) {
	var ret = ListRolesResponse{}
	return ret, c.requestV2(requestV2Opts{
		Method: Get,
		Path:   ListRolesPath,
		Ret:    &ret,
	})
}
