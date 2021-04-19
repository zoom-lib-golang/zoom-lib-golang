package zoom

import "fmt"

// GetAccountTrustedDomainsPath - v2 path for getting a specific user
const GetAccountTrustedDomainsPath = "/accounts/%s/trusted_domains"

// GetAccountTrustedDomainsOpts contains options for GetAccountTrustedDomains
type GetAccountTrustedDomainsOpts struct {
	AccountID string `url:"-"`
}

// GetAccountTrustedDomains calls /users/{userId}, searching for a user by ID or email, using the default client
func GetAccountTrustedDomains(opts GetAccountTrustedDomainsOpts) (AccountTrustedDomains, error) {
	return defaultClient.GetAccountTrustedDomains(opts)
}

// GetAccountTrustedDomains calls /users/{userId}, searching for a user by ID or email, using a specific client
func (c *Client) GetAccountTrustedDomains(opts GetAccountTrustedDomainsOpts) (AccountTrustedDomains, error) {
	var ret = AccountTrustedDomains{}
	return ret, c.requestV2(requestV2Opts{
		Method:        Get,
		Path:          fmt.Sprintf(GetAccountTrustedDomainsPath, opts.AccountID),
		URLParameters: opts,
		Ret:           &ret,
	})
}
