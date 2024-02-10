package zoom

import (
	"fmt"
	"log"
	"net/http"
)

func (c *Client) addRequestAuth(req *http.Request, err error) (*http.Request, error) {
	if err != nil {
		return nil, err
	}

	var token string
	switch c.AuthStrategy {
	case JWTAuth:
		token, err = jwtToken(c.Key, c.Secret)
		if err != nil {
			return nil, fmt.Errorf("failed to configure jwt token: %w", err)
		}

	case ServerToServerOAuth2:
		token, err = c.oauth2Handler.Authorize()
		if err != nil {
			return nil, fmt.Errorf("failed to set up server-to-server oauth2 token: %w", err)
		}
	}

	if Debug {
		log.Println("Auth Token: " + token + " with auth strategy: " + string(c.AuthStrategy))
	}

	// set Authorization header
	req.Header.Add("Authorization", "Bearer "+token)

	return req, nil
}
