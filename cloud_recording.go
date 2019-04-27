package zoom // Use this file for /recording endpoints

import "fmt"

const (
	// ListAllRecordingsPath - v2 lists all recordings
	ListAllRecordingsPath = "/users/%s/recordings"
)

// RecordingsFile represents a recordings file object
type RecordingsFile struct {
	ID             string `json:"id"`
	MeetingID      string `json:"meeting_id"`
	RecordingStart string `json:"recording_start"`
	RecordingEnd   string `json:"recording_end"`
	FileType       string `json:"file_type"`
	FileSize       int64  `json:"file_size"`
	PlayURL        string `json:"play_url"`
	DownloadURL    string `json:"download_url"`
	Status         string `json:"status"`
	DeletedTime    string `json:"deleted_time"`
	RecordingType  string `json:"recording_type"`
}

// Meeting represents a zoom meeting object
type Meeting struct {
	UUID            string           `json:"uuid"`
	ID              int              `json:"id"`
	AccountID       string           `json:"account_id"`
	HostID          string           `json:"host_id"`
	Topic           string           `json:"topic"`
	StartTime       *Time            `json:"start_time"`
	TotalSize       int              `json:"total_size"`
	RecordingCount  int              `json:"recording_count"`
	RecordingsFiles []RecordingsFile `json:"recordings_files"`
}

// ListAllRecordingsResponse contains the response from a call to ListAllRecordings
type ListAllRecordingsResponse struct {
	PageCount     int       `json:"page_count"`
	From          string    `json:"from"`
	To            string    `json:"to"`
	TotalRecords  int       `json:"total_records"`
	PageNumber    int       `json:"page_number"`
	NextPageToken string    `json:"next_page_token"`
	Meetings      []Meeting `json:"meetings"`
}

// ListAllRecordingsOptions contains options for ListAllRecordings.
type ListAllRecordingsOptions struct {
	UserID        string `url:"-"`
	PageSize      *int   `url:"page_size,omitempty"`
	NextPageToken string `json:"next_page_token,omitempty"`
	Mc            string `json:"mc"`
	Trash         bool   `json:"trash"`
	From          string `json:"from"`
	To            string `json:"to"`
}

// ListAllRecordings calls /users/{user_id}/recordings endpoint
// and gets all cloud recordings for a user, using the default
// client.
func ListAllRecordings(opts ListAllRecordingsOptions) (ListAllRecordingsResponse, error) {
	return defaultClient.ListAllRecordings(opts)
}

// ListAllRecordings calls /users/{user_id}/recordings endpoint
// and gets all cloud recordings for a user, using the c client
func (c *Client) ListAllRecordings(opts ListAllRecordingsOptions) (ListAllRecordingsResponse, error) {
	var ret = ListAllRecordingsResponse{}
	return ret, c.requestV2(requestV2Opts{
		Method:        Get,
		Path:          fmt.Sprintf(ListAllRecordingsPath, opts.UserID),
		URLParameters: &opts,
		Ret:           &ret,
	})
}
