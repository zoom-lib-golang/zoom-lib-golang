package zoom // types for /user requests

import "time"

const (
	// Basic user type
	Basic UserType = 1

	// Pro user type
	Pro UserType = 2

	// Corporate user type
	Corporate UserType = 3
)

// UserType is one of a fixed number of possible user types
type UserType int

// String provides a string representation of user types
func (t UserType) String() (str string) {
	switch t {
	case Basic:
		str = "basic"
	case Pro:
		str = "pro"
	case Corporate:
		str = "corporate"
	}
	return
}

// User represents an account user
type User struct {
	Email                            string    `json:"email"`
	ID                               string    `json:"id"`
	AccountID                        string    `json:"account_id"`
	CreatedAt                        time.Time `json:"created_at"`
	FirstName                        string    `json:"first_name"`
	LastName                         string    `json:"last_name"`
	PicURL                           string    `json:"pic_url"`
	Type                             UserType  `json:"type"`
	DisableChat                      bool      `json:"disable_chat"`
	EnableE2eEncryption              bool      `json:"enable_e2e_encryption"`
	EnableSilentMode                 bool      `json:"enable_silent_mode"`
	DisableGroupHd                   bool      `json:"disable_group_hd"`
	DisableRecording                 bool      `json:"disable_recording"`
	EnableCmr                        bool      `json:"enable_cmr"`
	EnableAutoRecording              bool      `json:"enable_auto_recording"`
	EnableCloudAutoRecording         bool      `json:"enable_cloud_auto_recording"`
	Verified                         int       `json:"verified"`
	PMI                              int       `json:"pmi"`
	MeetingCapacity                  int       `json:"meeting_capacity"`
	EnableWebinar                    bool      `json:"enable_webinar"`
	WebinarCapacity                  int       `json:"webinar_capacity"`
	EnableLarge                      bool      `json:"enable_large"`
	LargeCapacity                    int       `json:"large_capacity"`
	DisableFeedback                  bool      `json:"disable_feedback"`
	DisableJbhReminder               bool      `json:"disable_jbh_reminder"`
	EnableBreakoutRoom               bool      `json:"enable_breakout_room"`
	EnablePolling                    bool      `json:"enable_polling"`
	EnableAnnotation                 bool      `json:"enable_annotation"`
	EnableShareDualCamera            bool      `json:"enable_share_dual_camera"`
	EnableFarEndCameraControl        bool      `json:"enable_far_end_camera_control"`
	DisablePrivateChat               bool      `json:"disable_private_chat"`
	EnableEnterExitChime             bool      `json:"enable_enter_exit_chime"`
	DisableCancelMeetingNotification bool      `json:"disable_cancel_meeting_notification"`
	EnableRemoteSupport              bool      `json:"enable_remote_support"`
	EnableFileTransfer               bool      `json:"enable_file_transfer"`
	EnableVirtualBackground          bool      `json:"enable_virtual_background"`
	EnableAttentionTracking          bool      `json:"enable_attention_tracking"`
	EnableWaitingRoom                bool      `json:"enable_waiting_room"`
	EnableClosedCaption              bool      `json:"enable_closed_caption"`
	EnableCoHost                     bool      `json:"enable_co_host"`
	EnableAutoSavingChats            bool      `json:"enable_auto_saving_chats"`
	EnablePhoneParticipantsPassword  bool      `json:"enable_phone_participants_password"`
	EnableAutoDeleteCmr              bool      `json:"enable_auto_delete_cmr"`
	AutoDeleteCmrDays                int       `json:"auto_delete_cmr_days"`
	Dept                             string    `json:"dept"`
	LastClientVersion                string    `json:"lastClientVersion"`
	LastLoginTime                    string    `json:"lastLoginTime"`
	Token                            string    `json:"token"`
	ZPK                              string    `json:"zpk"`
}
