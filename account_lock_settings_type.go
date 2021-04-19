package zoom

// AccountLockSettings represents account lock settings
type AccountLockSettings struct {
	ScheduleMeeting   interface{} `json:"schedule_meeting,omitempty"`
	InMeeting         interface{} `json:"in_meeting,omitempty"`
	EmailNotification interface{} `json:"email_notification,omitempty"`
	ZoomRooms         interface{} `json:"zoom_rooms,omitempty"`
	Recording         interface{} `json:"recording,omitempty"`
	Telephony         interface{} `json:"telephony,omitempty"`
	TSP               interface{} `json:"tsp,omitempty"`
	MeetingSecurity   interface{} `json:"meeting_security,omitempty"`
}
