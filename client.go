package zoom

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
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

// Client is responsible for making API requests
type Client struct {
	Key       string
	Secret    string
	Transport http.RoundTripper
	Timeout   time.Duration // set to value > 0 to enable a request timeout
	endpoint  string
}

// NewClient returns a new API client
func NewClient(apiKey string, apiSecret string) *Client {
	var uri = url.URL{
		Scheme: "https",
		Host:   apiURI,
		Path:   apiVersion,
	}

	return &Client{
		Key:      apiKey,
		Secret:   apiSecret,
		endpoint: uri.String(),
	}
}

func normalizeParams(parameters interface{}) interface{} {
	if parameters == nil {
		return nil
	}

	switch reflect.TypeOf(parameters).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(parameters)
		if s.Len() == 0 {
			return nil
		}

		return s.Index(0).Interface()
	}

	return parameters
}

type requestV2Opts struct {
	Client         *Client
	Method         HTTPMethod
	Path           string
	URLParameters  interface{}
	DataParameters interface{}
	Ret            interface{}
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

	return client.Do(req)
}

func (c *Client) addRequestAuth(req *http.Request, err error) (*http.Request, error) {
	if err != nil {
		return nil, err
	}

	// establish JWT token
	ss, err := jwtToken(c.Key, c.Secret)
	if err != nil {
		return nil, err
	}

	// set JWT Authorization header
	req.Header.Add("Authorization", "Bearer "+ss)

	return req, nil
}

func (c *Client) httpRequest(opts requestV2Opts) (*http.Request, error) {
	var buf bytes.Buffer

	// allow variadic parameters
	dataParams := normalizeParams(opts.DataParameters)

	// encode body parameters if any
	if err := json.NewEncoder(&buf).Encode(&dataParams); err != nil {
		return nil, err
	}

	// set URL parameters
	values, err := query.Values(normalizeParams(opts.URLParameters))
	if err != nil {
		return nil, err
	}

	// set request URL
	requestURL := c.endpoint + opts.Path + "?" + values.Encode()

	if Debug {
		log.Printf("Request URL: %s", requestURL)
		log.Printf("URL Parameters: %s", values.Encode())
	}

	// create HTTP request
	req, err := http.NewRequest(string(opts.Method), requestURL, &buf)
	if err != nil {
		return nil, err
	}

	return req, nil
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

func request(c *Client, path string, parameters interface{}, ret interface{}) error {
	return errors.New("this method is DEPRECATED! Use requestV2. This method will be removed when v2 implementation is complete")
}