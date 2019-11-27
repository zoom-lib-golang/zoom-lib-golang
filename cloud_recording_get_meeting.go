package zoom

import "fmt"

const (
	// GetMeetingRecordingsPath - v2 get all the recordings from a meeting
	GetMeetingRecordingsPath = "/meetings/%s/recordings"
)

type (
	// GetMeetingRecordingsOptions contains options for GetMeetingRecordings
	GetMeetingRecordingsOptions struct {
		MeetingID string `url:"-"`
	}
)

// GetMeetingRecordings calls /meetings/{meetingId}/recordings endpoint
func GetMeetingRecordings(opts GetMeetingRecordingsOptions) (CloudRecordingMeeting, error) {
	return defaultClient.GetMeetingRecordings(opts)
}

// GetMeetingRecordings calls /meetings/{meetingId}/recordings endpoint
func (c *Client) GetMeetingRecordings(opts GetMeetingRecordingsOptions) (CloudRecordingMeeting, error) {
	var ret = CloudRecordingMeeting{}
	return ret, c.requestV2(requestV2Opts{
		Method:        Get,
		Path:          fmt.Sprintf(GetMeetingRecordingsPath, opts.MeetingID),
		URLParameters: &opts,
		Ret:           &ret,
	})
}
