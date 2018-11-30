package zoom

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/google/go-querystring/query"
	"gopkg.in/dgrijalva/jwt-go.v3"
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

func requestV2(c *Client, method Method, path string, urlParameters interface{}, dataParameters interface{}, ret interface{}) error {
	var buf bytes.Buffer

	// set client to default if not provided
	if c == nil {
		if defaultClient == nil {
			defaultClient = NewClient(APIKey, APISecret)
		}

		c = defaultClient
	}

	// allow variadic parameters
	urlParams := normalizeParams(urlParameters)
	dataParams := normalizeParams(dataParameters)

	// establish JWT token
	claims := &jwt.StandardClaims{
		Issuer:    c.Key,
		ExpiresAt: jwt.TimeFunc().Local().Add(time.Second * time.Duration(5000)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.Header["alg"] = "HS256"
	token.Header["typ"] = "JWT"
	ss, err := token.SignedString([]byte(c.Secret))
	if err != nil {
		return err
	}

	// set URL parameters
	values, err := query.Values(urlParams)
	if err != nil {
		return err
	}

	// set request URL
	requestURL := c.endpoint + path + "?"
	requestURL += values.Encode()

	if Debug {
		log.Printf("Request URL: %s", requestURL)
		log.Printf("URL Parameters: %s", values.Encode())
	}

	// create HTTP client and set timeout
	client := &http.Client{Transport: c.Transport}
	if c.Timeout > 0 {
		client.Timeout = c.Timeout
	}

	// encode body parameters if any
	if err := json.NewEncoder(&buf).Encode(&dataParams); err != nil {
		return err
	}

	// create HTTP request
	req, err := http.NewRequest(string(method), requestURL, &buf)
	if err != nil {
		return err
	}

	// set JWT Authorization header
	req.Header.Add("Authorization", "Bearer "+ss)

	// execute HTTP request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	// get HTTP response
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if Debug {
		log.Printf("Response Body: %s", string(body))
	}

	if err := checkError(body); err != nil {
		return err
	}

	return json.Unmarshal(body, &ret)
}

func request(c *Client, path string, parameters interface{}, ret interface{}) error {
	if c == nil {
		if defaultClient == nil {
			defaultClient = NewClient(APIKey, APISecret)
		}

		c = defaultClient
	}

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
		return err
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
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if Debug {
		log.Printf("Response Body: %s", string(body))
	}

	if err := checkError(body); err != nil {
		return err
	}

	return json.Unmarshal(body, &ret)
}
