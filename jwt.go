package zoom

import (
	"log"
	"net/http"
	"time"

	"gopkg.in/dgrijalva/jwt-go.v3"
)

func jwtToken(key string, secret string) (string, error) {
	claims := &jwt.StandardClaims{
		Issuer:    key,
		ExpiresAt: jwt.TimeFunc().Local().Add(time.Second * time.Duration(5000)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.Header["alg"] = "HS256"
	token.Header["typ"] = "JWT"
	return token.SignedString([]byte(secret))
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

	if Debug {
		log.Println("JWT Token: " + ss)
	}

	// set JWT Authorization header
	req.Header.Add("Authorization", "Bearer "+ss)

	return req, nil
}
