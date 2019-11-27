package zoom // Use this file for /recording endpoints

import "fmt"

const (
	// ListAllRecordingsPath - v2 lists all recordings
	ListAllRecordingsPath = "/users/%s/recordings"

	// TrashTypeMeetingRecordings list all meeting recordings from the trash. Default.
	TrashTypeMeetingRecordings TrashType = "meeting_recordings"
	// TrashTypeRecordingFile list all individual recording files from the trash
	TrashTypeRecordingFile TrashType = "recording_file"

	// RecordingTypeSharedScreenWithSpeakerViewCC is a shared screen with spearker view (CC) recording
	RecordingTypeSharedScreenWithSpeakerViewCC RecordingType = "shared_screen_with_speaker_view(CC)"
	// RecordingTypeSharedScreenWithSpeakerView is a shared screen with spearker view recording
	RecordingTypeSharedScreenWithSpeakerView RecordingType = "shared_screen_with_speaker_view"
	// RecordingTypeSharedScreenWithGalleryView is a shared screen with gallery view recording
	RecordingTypeSharedScreenWithGalleryView RecordingType = "shared_screen_with_gallery_view"
	// RecordingTypeSpeakerView is a speaker view recording
	RecordingTypeSpeakerView RecordingType = "speaker_view"
	// RecordingTypeGalleryView is a gallery view recording
	RecordingTypeGalleryView RecordingType = "gallery_view"
	// RecordingTypeSharedScreen is a shared screen recording
	RecordingTypeSharedScreen RecordingType = "shared_screen"
	// RecordingTypeAudioOnly is an audio only recording
	RecordingTypeAudioOnly RecordingType = "audio_only"
	// RecordingTypeAudioTranscript is an audio transcript recording
	RecordingTypeAudioTranscript RecordingType = "audio_transcript"
	// RecordingTypeChatFile is a chat file recording
	RecordingTypeChatFile RecordingType = "chat_file"
	// RecordingTypeTIMELINE is a timeline recording
	RecordingTypeTIMELINE RecordingType = "TIMELINE"
)

type (
	// TrashType is the type of Cloud recording that you would like to retrieve from the trash
	TrashType string

	// RecordingType is the recording file type
	RecordingType string

	// RecordingFile represents a recordings file object
	RecordingFile struct {
		ID             string `json:"id"`
		MeetingID      string `json:"meeting_id"`
		RecordingStart *Time  `json:"recording_start"`
		RecordingEnd   *Time  `json:"recording_end"`
		FileType       string `json:"file_type"`
		FileSize       int    `json:"file_size"`
		PlayURL        string `json:"play_url"`
		// The URL using which the recording file can be downloaded. To access a private or
		// password protected cloud recording, you must use a [Zoom JWT App Type]
		DownloadURL   string        `json:"download_url"`
		Status        string        `json:"status"`
		DeletedTime   *Time         `json:"deleted_time"`
		RecordingType RecordingType `json:"recording_type"`
	}

	// CloudRecordingMeeting represents a zoom meeting object
	CloudRecordingMeeting struct {
		UUID           string          `json:"uuid"`
		ID             string          `json:"id"`
		AccountID      string          `json:"account_id"`
		HostID         string          `json:"host_id"`
		Topic          string          `json:"topic"`
		StartTime      *Time           `json:"start_time"`
		Duration       int             `json:"duration"`
		TotalSize      string          `json:"total_size"`
		RecordingCount string          `json:"recording_count"`
		RecordingFiles []RecordingFile `json:"recording_files"`
	}

	// ListAllRecordingsResponse contains the response from a call to ListAllRecordings
	ListAllRecordingsResponse struct {
		From          *Date                   `json:"from"`
		To            *Date                   `json:"to"`
		PageCount     int                     `json:"page_count"`
		PageSize      int                     `json:"page_size"`
		TotalRecords  int                     `json:"total_records"`
		NextPageToken string                  `json:"next_page_token"`
		Meetings      []CloudRecordingMeeting `json:"meetings"`
	}

	// ListAllRecordingsOptions contains options for ListAllRecordings.
	ListAllRecordingsOptions struct {
		UserID        string `url:"-"`
		PageSize      *int   `url:"page_size,omitempty"`
		NextPageToken string `url:"next_page_token,omitempty"`
		Mc            string `url:"mc,omitempty"`
		Trash         bool   `url:"trash,omitempty"`
		// From is a YYYY-MM-DD string representing a date
		From string `url:"from"`
		// To is a YYYY-MM-DD string representing a date
		To        string    `url:"to"`
		TrashType TrashType `url:"trash_type,omitempty"`
	}
)

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
