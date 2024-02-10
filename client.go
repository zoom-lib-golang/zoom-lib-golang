package zoom

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/google/go-querystring/query"
)

const (
	apiURI     = "api.zoom.us"
	apiVersion = "/v2"
)

var (
	// Debug causes debugging message to be printed, using the log package,
	// when set to true
	Debug = false

	// APIKey is a package-wide API key, used when no client is instantiated
	APIKey string

	// APISecret is a package-wide API secret, used when no client is instantiated
	APISecret string

	defaultClient *Client
)

type AuthStrategy string

var (
	JWTAuth              = AuthStrategy("jwt")
	ServerToServerOAuth2 = AuthStrategy("server-to-server-oauth2")

	zoomAPIUrl = url.URL{
		Scheme: "https",
		Host:   apiURI,
		Path:   apiVersion,
	}
)

// Client is responsible for making API requests
type Client struct {
	AuthStrategy AuthStrategy

	Key    string
	Secret string

	oauth2Handler *oauth2Handler

	Transport http.RoundTripper
	Timeout   time.Duration // set to value > 0 to enable a request timeout
	endpoint  string
}

// NewClient returns a new API client
func NewClient(apiKey string, apiSecret string) *Client {
	return &Client{
		AuthStrategy: JWTAuth,
		Key:          apiKey,
		Secret:       apiSecret,
		endpoint:     zoomAPIUrl.String(),
	}
}

func NewOAuth2Client(accountID, clientID, ClientSecret string) *Client {

	handler := &oauth2Handler{
		oauthAccountID:    accountID,
		oauthClientID:     clientID,
		oauthClientSecret: ClientSecret,
	}
	_, err := handler.Authorize()
	if err != nil {
		panic(fmt.Sprintf("oauth2 client is not setup for success: %s", err))
	}

	return &Client{
		AuthStrategy:  ServerToServerOAuth2,
		oauth2Handler: handler,

		endpoint: zoomAPIUrl.String(),
	}

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
			defaultClient = NewClient(APIKey, APISecret)
		}

		return defaultClient
	}

	return c
}

func (c *Client) executeRequest(opts requestV2Opts) (*http.Response, error) {
	client := c.httpClient()
	req, err := c.addRequestAuth(c.httpRequest(opts))
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

	client := &http.Client{Transport: c.Transport}
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
