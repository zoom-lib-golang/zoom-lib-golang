package zoom

import "fmt"

// DeleteMeetingOptions are the options to delete a meeting
type DeleteMeetingOptions struct {
	MeetingID    int    `url:"-"`
	OccurrenceID string `url:"occurrence_id,omitempty"`
	// ScheduleForReminder notify host and alternative host about meeting cancellation via
	// email
	ScheduleForReminder bool `url:"schedule_for_reminder,omitempty"`
}

// DeleteMeetingPath - v2 delete a meeting
const DeleteMeetingPath = "/meetings/%d"

// DeleteMeeting calls DELETE /meetings/{meetingID}
func DeleteMeeting(opts DeleteMeetingOptions) error {
	return defaultClient.DeleteMeeting(opts)
}

// DeleteMeeting calls DELETE /meetings/{meetingID}
func (c *Client) DeleteMeeting(opts DeleteMeetingOptions) error {
	return c.requestV2(requestV2Opts{
		Method:        Delete,
		Path:          fmt.Sprintf(DeleteMeetingPath, opts.MeetingID),
		URLParameters: &opts,
		HeadResponse:  true,
	})
}
