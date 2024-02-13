package zoom

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/google/go-querystring/query"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const (
	apiURL     = "https://api.zoom.us/v2"
	tokenURL   = "https://zoom.us/oauth/token"
)

var (
	// Debug causes debugging message to be printed, using the log package,
	// when set to true
	Debug = false

	// Package-wide account ID
	AccountID string

	// ClientID is a package-wide client key, used when no client is instantiated
	ClientID string

	// ClientSecret is a package-wide client secret, used when no client is instantiated
	ClientSecret string

	defaultClient *Client
)

// Client is responsible for making API requests
type Client struct {
	clientcredentials.Config
	Timeout   time.Duration // set to value > 0 to enable a request timeout
	endpoint  string
}

// NewClient returns a new API client
func NewClient(accountID, clientID, clientSecret string, scopes []string) (client *Client) {
	params := url.Values{}
	params.Add("grant_type", "account_credentials")
	params.Add("account_id", accountID)

	client = &Client{
		Config: clientcredentials.Config{
			ClientID:       clientID,
			ClientSecret:   clientSecret,
			TokenURL:       tokenURL,
			Scopes:         scopes,
			EndpointParams: params,
			AuthStyle:      oauth2.AuthStyleInHeader,
		},
		endpoint: apiURL,
	}

	return
}

type requestV2Opts struct {
	Client         *Client
	Method         HTTPMethod
	URLParameters  interface{}
	Path           string
	DataParameters interface{}
	Ret            interface{}
	// HeadResponse represents responses that don't have a body
	HeadResponse bool
}

func initializeDefault(c *Client) *Client {
	if c == nil {
		if defaultClient == nil {
			defaultClient = NewClient(AccountID, ClientID, ClientSecret, nil)
		}

		return defaultClient
	}

	return c
}

func (c *Client) executeRequest(opts requestV2Opts) (*http.Response, error) {
	client := c.httpClient()
	req, err := c.httpRequest(opts)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	return client.Do(req)
}

func (c *Client) httpRequest(opts requestV2Opts) (*http.Request, error) {
	var buf bytes.Buffer

	// encode body parameters if any
	if err := json.NewEncoder(&buf).Encode(&opts.DataParameters); err != nil {
		return nil, err
	}

	// set URL parameters
	values, err := query.Values(opts.URLParameters)
	if err != nil {
		return nil, err
	}

	// set request URL
	requestURL := c.endpoint + opts.Path
	if len(values) > 0 {
		requestURL += "?" + values.Encode()
	}

	if Debug {
		log.Printf("Request URL: %s", requestURL)
		log.Printf("URL Parameters: %s", values.Encode())
		log.Printf("Body Parameters: %s", buf.String())
	}

	// create HTTP request
	return http.NewRequest(string(opts.Method), requestURL, &buf)
}

func (c *Client) httpClient() *http.Client {
	client := c.Config.Client(context.TODO())
	if c.Timeout > 0 {
		client.Timeout = c.Timeout
	}

	return client
}

func (c *Client) requestV2(opts requestV2Opts) error {
	// make sure the defaultClient is not nil if we are using it
	c = initializeDefault(c)

	// execute HTTP request
	resp, err := c.executeRequest(opts)
	if err != nil {
		return err
	}

	// If there is no body in response
	if opts.HeadResponse {
		return c.requestV2HeadOnly(resp)
	}

	return c.requestV2WithBody(opts, resp)
}

func (c *Client) requestV2WithBody(opts requestV2Opts, resp *http.Response) error {
	// read HTTP response
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if Debug {
		log.Printf("Response Body: %s", string(body))
	}

	// check for Zoom errors in the response
	if err := checkError(body); err != nil {
		return err
	}

	// unmarshal the response body into the return object
	return json.Unmarshal(body, &opts.Ret)
}

func (c *Client) requestV2HeadOnly(resp *http.Response) error {
	if resp.StatusCode != 204 {
		return errors.New(resp.Status)
	}

	// there were no errors, just return
	return nil
}
