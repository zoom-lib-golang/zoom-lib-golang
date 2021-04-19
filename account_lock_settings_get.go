package zoom

import "fmt"

// GetAccountLockSettingsPath - v2 path for getting a specific user
const GetAccountLockSettingsPath = "/accounts/%s/lock_settings"

// GetAccountLockSettingsOpts contains options for GetAccountLockSettings
type GetAccountLockSettingsOpts struct {
	AccountID         string `url:"-"`
	Option            string `url:"option,omitempty"`
	CustomQueryFields string `url:"custom_query_fields,omitempty"`
}

// GetAccountLockSettings calls /users/{userId}, searching for a user by ID or email, using the default client
func GetAccountLockSettings(opts GetAccountLockSettingsOpts) (AccountLockSettings, error) {
	return defaultClient.GetAccountLockSettings(opts)
}

// GetAccountLockSettings calls /users/{userId}, searching for a user by ID or email, using a specific client
func (c *Client) GetAccountLockSettings(opts GetAccountLockSettingsOpts) (AccountLockSettings, error) {
	var ret = AccountLockSettings{}
	return ret, c.requestV2(requestV2Opts{
		Method:        Get,
		Path:          fmt.Sprintf(GetAccountLockSettingsPath, opts.AccountID),
		URLParameters: opts,
		Ret:           &ret,
	})
}
