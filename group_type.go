package zoom

// Group represents an account group
type Group struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	TotalMembers int    `json:"total_members"`
}
