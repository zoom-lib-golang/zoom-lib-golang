package zoom

// AccountManagedDomains represents account managed domains
type AccountManagedDomains struct {
	TotalRecords int      `json:"total_records"`
	Domains      []string `json:"domains"`
}
