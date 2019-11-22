package zoom

import "fmt"

// CreateMeetingOptions are the options to create a meeting with
type CreateMeetingOptions struct {
	HostID         string          `json:"-"`
	Topic          string          `json:"topic,omitempty"`
	Type           MeetingType     `json:"type,omitempty"`
	StartTime      *Time           `json:"start_time,omitempty"`
	Duration       int             `json:"duration,omitempty"`
	Timezone       string          `json:"timezone,omitempty"`
	Password       string          `json:"password,omitempty"` // Max 10 characters. [a-z A-Z 0-9 @ - _ *]
	Agenda         string          `json:"agenda,omitempty"`
	TrackingFields []TrackingField `json:"tracking_fields,omitempty"`
	Settings       MeetingSettings `json:"settings,omitempty"`
}

// CreateMeetingPath - v2 create a meeting for a user
const CreateMeetingPath = "/users/%s/meetings"

// CreateMeeting calls POST /users/{userId}/meetings
func CreateMeeting(opts CreateMeetingOptions) (Meeting, error) {
	return defaultClient.CreateMeeting(opts)
}

// CreateMeeting calls POST /users/{userId}/meetings
// https://marketplace.zoom.us/docs/api-reference/zoom-api/meetings/meetingcreate
func (c *Client) CreateMeeting(opts CreateMeetingOptions) (Meeting, error) {
	var ret = Meeting{}
	return ret, c.requestV2(requestV2Opts{
		Method:         Post,
		Path:           fmt.Sprintf(CreateMeetingPath, opts.HostID),
		DataParameters: &opts,
		Ret:            &ret,
	})
}
