package zoom

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

const zoomOAuth2TokenURL = "https://zoom.us/oauth/token"

type oauth2Handler struct {
	token *oauth2Token

	oauthAccountID    string
	oauthClientID     string
	oauthClientSecret string
}

type oauth2Token struct {
	Token     string
	ExpiresAt time.Time
}

func (t oauth2Token) isValid() bool {
	if t.Token == "" { // Empty token is always invalid
		return false
	}

	return time.Now().Before(t.ExpiresAt)
}

type zoomReply struct {
	AccessToken      string `json:"access_token"`
	ExpiresInSeconds int    `json:"expires_in"`
}

func (o *oauth2Handler) Authorize() (string, error) {

	if o.token != nil && o.token.isValid() {
		return o.token.Token, nil // reuse tokens while they are valid
	}

	uri, err := url.Parse(zoomOAuth2TokenURL)
	if err != nil {
		return "", fmt.Errorf("could not parse URL %s to authorize oauth2: %w", zoomOAuth2TokenURL, err)
	}

	q := uri.Query()
	q.Add("grant_type", "account_credentials")
	q.Add("account_id", o.oauthAccountID)
	uri.RawQuery = q.Encode()

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, uri.String(), nil)
	if err != nil {
		return "", fmt.Errorf("could not create request to authorize OAuth2: %w", err)
	}
	req.SetBasicAuth(o.oauthClientID, o.oauthClientSecret)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to execute request to authorize oauth2 server-to-server token: %w", err)
	}

	b, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read body when authorizing oauth2 server-to-server token: %w", err)
	}

	switch response.StatusCode {
	case http.StatusOK:

		if Debug {
			log.Printf("Got json reply: %s\n", string(b))
		}

		reply := zoomReply{}
		if err := json.Unmarshal(b, &reply); err != nil {
			return "", fmt.Errorf("failed to unmarshal json body when authorizing oauth2 server-to-server token: %w", err)
		}

		if reply.AccessToken == "" {
			return "", fmt.Errorf("json reply when authorizing oauth2 server-to-server returned an empty access token: %s", string(b))
		}

		o.token = &oauth2Token{
			Token:     reply.AccessToken,
			ExpiresAt: time.Now().Add(time.Duration(reply.ExpiresInSeconds) * time.Second),
		}

		return o.token.Token, nil

	default:
		return "", fmt.Errorf("failed to authorize oauth2 server-to-server token, got status code %d with body %s on request %+v", response.StatusCode, string(b), req)
	}

}
