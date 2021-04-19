package zoom

import "fmt"

// GetAccountSettingsPath - path to get account settings
const GetAccountSettingsPath = "/accounts/%s/settings"

// GetAccountSettingsOpts contains options for GetAccountSettings
type GetAccountSettingsOpts struct {
	AccountID         string `url:"-"`
	Option            string `url:"option,omitempty"`
	CustomQueryFields string `url:"custom_query_fields,omitempty"`
}

// GetAccountSettings gets the account settings for a master or sub account
func GetAccountSettings(opts GetAccountSettingsOpts) (AccountSettings, error) {
	return defaultClient.GetAccountSettings(opts)
}

// GetAccountSettings gets the account settings for a master or sub account
func (c *Client) GetAccountSettings(opts GetAccountSettingsOpts) (AccountSettings, error) {
	var ret = AccountSettings{}
	return ret, c.requestV2(requestV2Opts{
		Method:        Get,
		Path:          fmt.Sprintf(GetAccountSettingsPath, opts.AccountID),
		URLParameters: opts,
		Ret:           &ret,
	})
}
