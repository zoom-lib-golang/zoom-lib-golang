package zoom

import "fmt"

// GetAccountManagedDomainsPath - v2 path for getting a specific user
const GetAccountManagedDomainsPath = "/accounts/%s/managed_domains"

// GetAccountManagedDomainsOpts contains options for GetAccountManagedDomains
type GetAccountManagedDomainsOpts struct {
	AccountID string `url:"-"`
}

// GetAccountManagedDomains calls /users/{userId}, searching for a user by ID or email, using the default client
func GetAccountManagedDomains(opts GetAccountManagedDomainsOpts) (AccountManagedDomains, error) {
	return defaultClient.GetAccountManagedDomains(opts)
}

// GetAccountManagedDomains calls /users/{userId}, searching for a user by ID or email, using a specific client
func (c *Client) GetAccountManagedDomains(opts GetAccountManagedDomainsOpts) (AccountManagedDomains, error) {
	var ret = AccountManagedDomains{}
	return ret, c.requestV2(requestV2Opts{
		Method:        Get,
		Path:          fmt.Sprintf(GetAccountManagedDomainsPath, opts.AccountID),
		URLParameters: opts,
		Ret:           &ret,
	})
}
