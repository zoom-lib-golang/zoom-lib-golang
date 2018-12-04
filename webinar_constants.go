package zoom

// WebinarType is one of a fixed number of possible webinar types
type WebinarType int

const (
	// TypeWebinar is the default webinar type
	TypeWebinar WebinarType = 5

	// TypeRecurringWebinar is a recurring webinar
	TypeRecurringWebinar WebinarType = 6

	// TypeRecurringWebinarFixedTime is a recurring webinar with fixed time
	TypeRecurringWebinarFixedTime WebinarType = 9
)

// PurchasingTimeFrameType contains the list of valid values for purchasing
// timeframe
type PurchasingTimeFrameType string

const (
	// WithinMonth is "Within a month" purchasing timeframe
	WithinMonth PurchasingTimeFrameType = "Within a month"

	// OneToThreeMonths is "1-3 months" purchasing timeframe
	OneToThreeMonths PurchasingTimeFrameType = "1-3 months"

	// FourToSixMonths is "4-6 months" purchasing timeframe
	FourToSixMonths PurchasingTimeFrameType = "4-6 months"

	// MoreThan6Months is "More than 6 months" purchasing timeframe
	MoreThan6Months PurchasingTimeFrameType = "More than 6 months"

	// NoTimeframe is no purchasing timeframe
	NoTimeframe PurchasingTimeFrameType = "No timeframe"
)

// PurchaseProcessRoleType contains a lsit of valid values for a registrant's
// role in the purchasing process
type PurchaseProcessRoleType string

const (
	// DecisionMaker role in purchasing
	DecisionMaker PurchaseProcessRoleType = "Decision Maker"

	// EvaluatorRecommender role in purchasing
	EvaluatorRecommender PurchaseProcessRoleType = "Evaluator/Recommender"

	// Influencer role in purchasing
	Influencer PurchaseProcessRoleType = "Influencer"

	// NotInvolved has no role in purchasing
	NotInvolved PurchaseProcessRoleType = "Not involved"
)

// NumberOfEmployeesType contains a list of valid values for number of
// employees in a registrant's company
type NumberOfEmployeesType string

const (
	// OneToTwenty is "1-20" employees
	OneToTwenty NumberOfEmployeesType = "1-20"

	// TwentyOneToFifty is "21-50" employees
	TwentyOneToFifty NumberOfEmployeesType = "21-50"

	// FiftyOneToOneHundred is "51-100" employees
	FiftyOneToOneHundred NumberOfEmployeesType = "51-100"

	// OneHundredOneToTwoFifty is "101-250" employees
	OneHundredOneToTwoFifty NumberOfEmployeesType = "101-250"

	// TwoFiftyOneToFiveHundred is "251-500" employees
	TwoFiftyOneToFiveHundred NumberOfEmployeesType = "251-500"

	// FiveHundredOneToOneThousand is "501-1,000" employees
	FiveHundredOneToOneThousand NumberOfEmployeesType = "501-1,000"

	// OneThousdandOneToFiveThousand is "1,001-5,001" employees
	OneThousdandOneToFiveThousand NumberOfEmployeesType = "1,001-5,001"

	// FiveThousandOneToTenThousand is "5,001-10,000" employees
	FiveThousandOneToTenThousand NumberOfEmployeesType = "5,001-10,000"

	// MoreThanTenThousand is "More than 10,000" employees
	MoreThanTenThousand NumberOfEmployeesType = "More than 10,000"
)

// ListWebinarRegistrantsStatusType contains possible options for "status" field when
// listing registrants
type ListWebinarRegistrantsStatusType string

const (
	// PendingApprovalType - registrants pending approval
	PendingApprovalType ListWebinarRegistrantsStatusType = "pending"

	// ApprovedType - approved registrants
	ApprovedType ListWebinarRegistrantsStatusType = "approved"

	// DeniedType - denied registrants
	DeniedType ListWebinarRegistrantsStatusType = "denied"
)
