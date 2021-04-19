package zoom

// AccountSettings represents account settings
type AccountSettings struct {
	ScheduleMeeting         interface{} `json:"schedule_meeting,omitempty"`
	InMeeting               interface{} `json:"in_meeting,omitempty"`
	EmailNotification       interface{} `json:"email_notification,omitempty"`
	ZoomRooms               interface{} `json:"zoom_rooms,omitempty"`
	Security                interface{} `json:"security,omitempty"`
	Recording               interface{} `json:"recording,omitempty"`
	Telephony               interface{} `json:"telephony,omitempty"`
	TSP                     interface{} `json:"tsp,omitempty"`
	Integration             interface{} `json:"integration,omitempty"`
	Feature                 interface{} `json:"feature,omitempty"`
	MeetingAuthentication   bool        `json:"meeting_authentication,omitempty"`
	RecordingAuthentication bool        `json:"recording_authentication,omitempty"`
	AuthenticationOptions   interface{} `json:"authentication_options,omitempty"`
	MeetingSecurity         interface{} `json:"meeting_security,omitempty"`
}
