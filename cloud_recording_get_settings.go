package zoom

import "fmt"

const (
	// GetMeetingRecordingSettingsPath - get the settings for a recording
	GetMeetingRecordingSettingsPath = "/meetings/%d/recordings/settings"
)

type (
	// GetMeetingRecordingSettingsOptions contains options for GetMeetingRecordingSettings
	GetMeetingRecordingSettingsOptions struct {
		MeetingID int `url:"-"`
	}
)

// GetMeetingRecordingSettings calls /meetings/{meetingId}/recordings/settings endpoint
func GetMeetingRecordingSettings(opts GetMeetingRecordingSettingsOptions) (CloudRecordingSettings, error) {
	return defaultClient.GetMeetingRecordingSettings(opts)
}

// GetMeetingRecordingSettings calls /meetings/{meetingId}/recordings/settings endpoint
func (c *Client) GetMeetingRecordingSettings(opts GetMeetingRecordingSettingsOptions) (CloudRecordingSettings, error) {
	var ret = CloudRecordingSettings{}
	return ret, c.requestV2(requestV2Opts{
		Method:        Get,
		Path:          fmt.Sprintf(GetMeetingRecordingSettingsPath, opts.MeetingID),
		URLParameters: &opts,
		Ret:           &ret,
	})
}

// CloudRecordingSettings represents settings for a recording
type CloudRecordingSettings struct {
	ShareRecording          string `json:"share_recording"`
	RecordingAuthentication bool   `json:"recording_authentication"`
	AuthenticationOption    string `json:"authentication_option"`
	AuthenticationDomains   string `json:"authentication_domains"`
	ViewerDownload          bool   `json:"viewer_download"`
	Password                string `json:"password"`
	OnDemand                bool   `json:"on_demand"`
	ApprovalType            int    `json:"approval_type"`
	SendEmailToHost         bool   `json:"send_email_to_host"`
	ShowSocialShareButtons  bool   `json:"show_social_share_buttons"`
	Topic                   string `json:"topic"`
}
