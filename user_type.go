package zoom // types for /user endpoints

// CreateUserAction specifies how to create a new user
type CreateUserAction string

// DeleteUserAction specifies how to delete a new user
type DeleteUserAction string

// UserType is one of a fixed number of possible user types
type UserType int

// UserLoginType is one of a fixed number of possible user login type
type UserLoginType int

// UserStatus is a user's active status, for ListUser
type UserStatus string

const (
	// Create action will be send new user a confirmation email required to activate
	Create CreateUserAction = "create"

	// AutoCreate action does not send the user a confirmation email
	AutoCreate CreateUserAction = "autoCreate"

	// CustCreate action creates a user without a password that cannot log into Zoom web portal or Zoom client
	CustCreate CreateUserAction = "custCreate"

	// SSOCreate action is provided for when the "Pre-Provisioning SSO User" option is enabled
	SSOCreate CreateUserAction = "ssoCreate"

	// DisassociateAction action disassociates a user
	DisassociateAction DeleteUserAction = "disassociate"

	// DeleteAction action deletes a user
	DeleteAction DeleteUserAction = "delete"

	// Basic user type
	Basic UserType = 1

	// Licensed user type
	Licensed UserType = 2

	// OnPrem user type
	OnPrem UserType = 3

	// Facebook user login type
	Facebook UserLoginType = 0

	// Google user login type
	Google UserLoginType = 1

	// API user login type
	API UserLoginType = 99

	// Zoom user login type
	Zoom UserLoginType = 100

	// SSO single sign on user login type
	SSO UserLoginType = 101

	// Active status
	Active UserStatus = "active"

	// Inactive status
	Inactive UserStatus = "inactive"

	// Pending status
	Pending UserStatus = "pending"
)

// String provides a string representation of user types
func (t UserType) String() (str string) {
	switch t {
	case Basic:
		str = "basic"
	case Licensed:
		str = "licensed"
	case OnPrem:
		str = "onPrem"
	}
	return
}

// User represents an account user
type User struct {
	// Fields returned in both List & Get
	ID                string   `json:"id"`
	FirstName         string   `json:"first_name"`
	LastName          string   `json:"last_name"`
	Email             string   `json:"email"`
	Type              int      `json:"type"`
	Status            string   `json:"status"`
	PMI               int      `json:"pmi"`
	Timezone          string   `json:"timezone"`
	Dept              string   `json:"dept"`
	RoleID            string   `json:"role_id"`
	CreatedAt         Time     `json:"created_at,string"`
	LastLoginTime     Time     `json:"last_login_time,string"`
	LastClientVersion string   `json:"last_client_version"`
	GroupIDs          []string `json:"group_ids"`
	IMGroupIDs        []string `json:"im_group_ids"`
	Verified          int      `json:"verified"`
	// Returned only when specifically requested via include_fields in the List call
	CustomAttributes []CustomAttribute `json:"custom_attributes"`
	HostKey          string            `json:"host_key"`
	// Returned by the Get call only
	PlanUnitedType     string        `json:"plan_united_type"`
	UsePMI             bool          `json:"use_pmi"`
	Language           string        `json:"language"`
	PhoneNumbers       []PhoneNumber `json:"phone_numbers"`
	VanityURL          string        `json:"vanity_url"`
	PersonalMeetingURL string        `json:"personal_meeting_url"`
	PicURL             string        `json:"pic_url"`
	CMSUserID          string        `json:"cms_user_id"`
	AccountID          string        `json:"account_id"`
	JID                string        `json:"jid"`
	JobTitle           string        `json:"job_title"`
	Company            string        `json:"company"`
	Location           string        `json:"location"`
	LoginType          int           `json:"login_type"`
}

// CustomAttribute represents a custom attribute
type CustomAttribute struct {
	Key   string `json:"key"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

// PhonNumbers represents represents a custom attribute
type PhoneNumber struct {
	Country  string `json:"country"`
	Code     string `json:"code"`
	Number   string `json:"number"`
	Verified bool   `json:"verified"`
}
