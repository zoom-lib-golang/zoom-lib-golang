package zoom

// Role represents an account role
type Role struct {
	ID                   string                   `json:"id"`
	Name                 string                   `json:"name"`
	Description          string                   `json:"description"`
	TotalMembers         int                      `json:"total_members"`
	Privileges           []string                 `json:"privileges"`
	SubAccountPrivileges RoleSubAccountPrivileges `json:"sub_account_privileges"`
}

// RoleSubAccountPrivileges are a partner plan feature
type RoleSubAccountPrivileges struct {
	SecondLevel int `json:"second_level"`
}
