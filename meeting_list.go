package zoom

import "fmt"

type (
	// ListMeetingType are the allowed meeting types
	ListMeetingType string

	// ListMeetingsOptions contains options for ListMeetings
	ListMeetingsOptions struct {
		HostID     string          `url:"-"`
		Type       ListMeetingType `url:"type,omitempty"`
		PageSize   *int            `url:"page_size,omitempty"`   // Default: 30, Maximum: 300
		PageNumber *int            `url:"page_number,omitempty"` // Default: 1
	}

	// ListMeetingsResponse container the response from a call to ListMeetings
	ListMeetingsResponse struct {
		PageCount    int           `json:"page_count"`
		TotalRecords int           `json:"total_records"`
		PageNumber   int           `json:"page_number"`
		PageSize     int           `json:"page_size"`
		Meetings     []ListMeeting `json:"meetings"`
	}

	// ListMeeting represents a meeting object returned by ListMeetings endpoint
	ListMeeting struct {
		UUID      string      `json:"uuid"`
		ID        int         `json:"id"`
		HostID    string      `json:"host_id"`
		Topic     string      `json:"topic"`
		Type      MeetingType `json:"type"`
		StartTime *Time       `json:"start_time"`
		Duration  int         `json:"duration"`
		Timezone  string      `json:"timezone"`
		CreatedAt *Time       `json:"created_at"`
		JoinURL   string      `json:"join_url"`
		Agenda    string      `json:"agenda"`
	}
)

const (
	// ListMeetingsPath - v2 lists all the meetings that were scheduled for a user
	ListMeetingsPath = "/users/%s/meetings"

	// ListMeetingTypeScheduled is a meeting that is scheduled
	ListMeetingTypeScheduled ListMeetingType = "scheduled"
	// ListMeetingTypeLive is a live meeting
	ListMeetingTypeLive ListMeetingType = "live" // DEFAULT
	// ListMeetingTypeUpcoming is an upcoming meeting
	ListMeetingTypeUpcoming ListMeetingType = "upcoming"
)

// ListMeetings calls /users/ID/meetings
func ListMeetings(opts ListMeetingsOptions) (ListMeetingsResponse, error) {
	return defaultClient.ListMeetings(opts)
}

// ListMeetings calls /users/ID/meetings
// https://marketplace.zoom.us/docs/api-reference/zoom-api/meetings/meetings
func (c *Client) ListMeetings(opts ListMeetingsOptions) (ListMeetingsResponse, error) {
	var ret = ListMeetingsResponse{}
	return ret, c.requestV2(requestV2Opts{
		Method:        Get,
		Path:          fmt.Sprintf(ListMeetingsPath, opts.HostID),
		URLParameters: &opts,
		Ret:           &ret,
	})
}
