package zoom

import "fmt"

// GetMeetingOptions are the options to get a meeting
type GetMeetingOptions struct {
	MeetingID    int    `url:"-"`
	OccurrenceID string `url:"occurrence_id,omitempty"`
}

// GetMeetingPath - v2 retrieve the details of a meeting
const GetMeetingPath = "/meetings/%d"

// GetMeeting calls /meetings/ID
func GetMeeting(opts GetMeetingOptions) (Meeting, error) {
	return defaultClient.GetMeeting(opts)
}

// GetMeeting calls /meetings/ID
// https://marketplace.zoom.us/docs/api-reference/zoom-api/meetings/meeting
func (c *Client) GetMeeting(opts GetMeetingOptions) (Meeting, error) {
	var ret = Meeting{}
	return ret, c.requestV2(requestV2Opts{
		Method:        Get,
		Path:          fmt.Sprintf(GetMeetingPath, opts.MeetingID),
		URLParameters: &opts,
		Ret:           &ret,
	})
}
