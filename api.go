package zoom

import (
	"encoding/json"
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
	apiVersion = "/v1"
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

type apiParams struct {
	APIKey    string `url:"api_key"`
	APISecret string `url:"api_secret"`
	DataType  string `url:"data_type"`
}

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

func request(c *Client, path string, parameters interface{}) ([]byte, error) {
	// allow variadic parameters
	switch reflect.TypeOf(parameters).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(parameters)
		if s.Len() == 0 {
			parameters = nil
		} else {
			parameters = s.Index(0).Interface()
		}
	}

	values, err := query.Values(parameters)
	if err != nil {
		return nil, err
	}

	values.Set("api_key", c.Key)
	values.Set("api_secret", c.Secret)
	values.Set("data_type", "JSON")

	requestURL := c.endpoint + path
	if Debug {
		log.Printf("Request URL: %s", requestURL)
		log.Printf("Request Parameters: %s", values.Encode())
	}

	client := &http.Client{Transport: c.Transport}
	if c.Timeout > 0 {
		client.Timeout = c.Timeout
	}

	// all v1 API endpoints use POST requests
	resp, err := client.PostForm(requestURL, values)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if Debug {
		log.Printf("Response Body: %s", string(body))
	}

	if err := checkError(body); err != nil {
		return nil, err
	}

	return body, nil
}

func executeRequest(c *Client, path string, parameters interface{}, ret interface{}) error {
	if c == nil {
		if defaultClient == nil {
			defaultClient = NewClient(APIKey, APISecret)
		}

		c = defaultClient
	}

	body, err := request(c, path, parameters)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, &ret)
}
