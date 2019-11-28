package zoom

const (
	// MeetingTypeInstant is an instant meeting
	MeetingTypeInstant MeetingType = 1
	// MeetingTypeScheduled is a scheduled meeting
	MeetingTypeScheduled MeetingType = 2 // Default
	// MeetingTypeRecurringNoFixedTime is a recurring meeting with no fixed time
	MeetingTypeRecurringNoFixedTime MeetingType = 3
	// MeetingTypeRecurring is a recurring meeting with fixed time
	MeetingTypeRecurring MeetingType = 8

	// MeetingStatusWaiting is a meeting that is waiting
	MeetingStatusWaiting MeetingStatus = "waiting"
	// MeetingStatusStarted is a meeting that is started
	MeetingStatusStarted MeetingStatus = "started"
	// MeetingStatusFinished is a meeting that is finished
	MeetingStatusFinished MeetingStatus = "finished"

	// ApprovalTypeAutomaticallyApprove is an automatically approved meeting
	ApprovalTypeAutomaticallyApprove ApprovalType = 0
	// ApprovalTypeManuallyApprove is a meeting that requires manual approval
	ApprovalTypeManuallyApprove ApprovalType = 1
	// ApprovalTypeNoRegistrationRequired is a meeting that requires no registration
	ApprovalTypeNoRegistrationRequired ApprovalType = 2 // DEFAULT

	// RegistrationTypeRegisterOnce Attendees register once and can attend any of the occurrences
	RegistrationTypeRegisterOnce RegistrationType = 1
	//RegistrationTypeRegisterEachTime  Attendeed need to register for each occurrence to attend
	RegistrationTypeRegisterEachTime RegistrationType = 2
	//RegistrationTypeRegisterOnceAndChooseOccurrences Attendees register once and can choose one or more occurrences to attend
	RegistrationTypeRegisterOnceAndChooseOccurrences RegistrationType = 3

	// AudioBoth is a meeting that allows telephony and VoIP
	AudioBoth Audio = "both"
	// AudioTelephony is a meeting that is telephony only
	AudioTelephony Audio = "telephony"
	// AudioVoIP is a meeting that is VoIP only
	AudioVoIP Audio = "voip"

	// AutoRecordingLocal record on local
	AutoRecordingLocal AutoRecording = "local"
	// AutoRecordingCloud record on cloud
	AutoRecordingCloud AutoRecording = "cloud"
	// AutoRecordingNone disabled
	AutoRecordingNone AutoRecording = "none"

	// GlobalDialInNumberTypeToll toll type of number
	GlobalDialInNumberTypeToll GlobalDialInNumberType = "toll"
	// GlobalDialInNumberTypeTollFree toll free type of number
	GlobalDialInNumberTypeTollFree GlobalDialInNumberType = "tollfree"

	// RecurrenceTypeDaily daily recurrence
	RecurrenceTypeDaily RecurrenceType = 1
	// RecurrenceTypeWeekly weekly recurrence
	RecurrenceTypeWeekly RecurrenceType = 2
	// RecurrenceTypeMonthly monthly recurrence
	RecurrenceTypeMonthly RecurrenceType = 3

	// MonthlyWeekLast last week of the month
	MonthlyWeekLast MonthlyWeek = -1
	// MonthlyWeekFirst first week of the month
	MonthlyWeekFirst MonthlyWeek = 1
	// MonthlyWeekSecond second week of the month
	MonthlyWeekSecond MonthlyWeek = 2
	// MonthlyWeekThird third week of the month
	MonthlyWeekThird MonthlyWeek = 3
	// MonthlyWeekFourth fourth week of the month
	MonthlyWeekFourth MonthlyWeek = 4
	//

	// WeekDaySunday Sunday
	WeekDaySunday WeekDay = 1
	// WeekDayMonday Monday
	WeekDayMonday WeekDay = 2
	// WeekDayTuesday Tuesday
	WeekDayTuesday WeekDay = 3
	// WeekDayWednesday Wednesday
	WeekDayWednesday WeekDay = 4
	// WeekDayThursday Thursday
	WeekDayThursday WeekDay = 5
	// WeekDayFriday Friday
	WeekDayFriday WeekDay = 6
	// WeekDaySaturday Saturday
	WeekDaySaturday WeekDay = 7
)

type (
	// MeetingType is the type of the meeting returned by the API
	MeetingType int

	// MeetingStatus is the status of the meeting
	MeetingStatus string

	// ApprovalType is the type of approval
	ApprovalType int

	// RegistrationType is the type of registration
	RegistrationType int

	// Audio determines how participants can join the audio portion of the meeting
	Audio string

	// AutoRecording automatic recording
	AutoRecording string

	// GlobalDialInNumberType is the type of global dial in number
	GlobalDialInNumberType string

	// RecurrenceType is the type of recurrence
	RecurrenceType int

	// MonthlyWeek the week a meeting will recur each month
	MonthlyWeek int

	// WeekDay is the day of the week
	WeekDay int

	// Meeting represents a meeting created/returned by GetMeeting
	Meeting struct {
		UUID              string        `json:"uuid,omitempty"`
		ID                int           `json:"id,omitempty"`
		HostID            string        `json:"host_id,omitempty"`
		Topic             string        `json:"topic"`
		Type              MeetingType   `json:"type"`
		Status            MeetingStatus `json:"status"`
		StartTime         *Time         `json:"start_time"`
		Duration          int           `json:"duration"`
		Timezone          string        `json:"timezone"`
		CreatedAt         *Time         `json:"created_at"`
		Agenda            string        `json:"agenda"`
		StartURL          string        `json:"start_url"`
		JoinURL           string        `json:"join_url"`
		Password          string        `json:"password"`
		H323Password      string        `json:"h323_password"`
		EncryptedPassword string        `json:"encrypted_password"`
		// PMI is Personal Meeting ID. Only used for scheduled meetings and recurring meetings with
		// no fixed time
		PMI            int64           `json:"pmi"`
		TrackingFields []TrackingField `json:"tracking_fields"`
		Occurrences    []Occurrence    `json:"occurrences"`
		Settings       MeetingSettings `json:"settings"`
		Recurrence     Recurrence      `json:"recurrence"`
	}

	// TrackingField is a tracking field
	TrackingField struct {
		Field string `json:"field"`
		Value string `json:"value"`
	}

	// Occurrence is an occurrence object
	Occurrence struct {
		ID        int    `json:"occurrence_id"`
		StartTime *Time  `json:"start_time"`
		Duration  int    `json:"duration"`
		Status    string `json:"status"`
	}

	// MeetingSettings are a meeting setting
	MeetingSettings struct {
		HostVideo        bool `json:"host_video,omitempty"`
		ParticipantVideo bool `json:"participant_video,omitempty"`
		// ChinaMeeting host meeting in China
		ChinaMeeting bool `json:"cn_meeting,omitempty"`
		// IndiaMeeting host meeting in India
		IndiaMeeting   bool `json:"in_meeting,omitempty"`
		JoinBeforeHost bool `json:"join_before_host,omitempty"`
		MuteUponEntry  bool `json:"mute_upon_entry,omitempty"`
		// Watermark add watermark when viewing a shared screen
		Watermark bool `json:"watermark,omitempty"`
		// Use Personal Meeting ID. Only used for scheduled meetings and recurring meetings with no
		// fixed time
		UsePMI           bool             `json:"use_pmi,omitempty"`
		ApprovalType     ApprovalType     `json:"approval_type,omitempty"`
		RegistrationType RegistrationType `json:"registration_type,omitempty"`
		// Audio determines how participants can join the audio portion of the meeting
		Audio         Audio         `json:"audio,omitempty"`
		AutoRecording AutoRecording `json:"auto_recording,omitempty"`
		// Only signed in users can join this meeting
		EnforceLogin bool `json:"enforce_login,omitempty"`
		// Only signed in users with specifid domains can join this meeting
		EnforceLoginDomains string `json:"enforce_login_domains,omitempty"`
		// Alternative host's emails or IDs: multiple values separated by comma
		AlternativeHosts string `json:"alternative_hosts,omitempty"`
		// CloseRegistration after event date
		CloseRegistration bool `json:"close_registration,omitempty"`
		// Enable waiting room
		WaitingRoom        bool                 `json:"waiting_room,omitempty"`
		GobalDialInNumbers []GlobalDialInNumber `json:"global_dial_in_numbers,omitempty"`
		ContactName        string               `json:"contact_name,omitempty"`
		ContactEmail       string               `json:"contact_email,omitempty"`
		// Send confirmation email to registrants
		RegistrantsConfirmationEmail bool `json:"registrants_confirmation_email,omitempty"`
	}

	// GlobalDialInNumber is a global dial in number
	GlobalDialInNumber struct {
		Country     string                 `json:"country"`
		CountryName string                 `json:"country_name"`
		City        string                 `json:"city"`
		Number      string                 `json:"number"`
		Type        GlobalDialInNumberType `json:"type"`
	}

	// Recurrence of the meeting
	Recurrence struct {
		Type           RecurrenceType `json:"type"`
		RepeatInterval int            `json:"repeat_interval"`
		WeeklyDays     string         `json:"weekly_days"`
		MonthlyDay     int            `json:"monthly_day"`
		MonthlyWeek    MonthlyWeek    `json:"monthly_week"`
		MonthlyWeekDay WeekDay        `json:"monthly_week_day"`
		// EndTimes how many times the meeting will recur before it is canceled (cannot be used
		// with "end_time_date"
		EndTimes int `json:"end_times"`
		// EndDateTime should be in UTC. Cannot be used with "end_times"
		EndDateTime *Time `json:"end_date_time"`
	}
)
