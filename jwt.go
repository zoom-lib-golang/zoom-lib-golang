package zoom

import (
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