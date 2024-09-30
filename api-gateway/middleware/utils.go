package middleware

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// ErrTokenNotFound ...
var (
	ErrTokenNotFound = errors.New("Missing or empty Authorization header, expected Bearer token")
	ErrNilRequest    = errors.New("Request nil")
)


func FromHeaderQuery(r *http.Request) (string, error) {

	raw := ""

	raw = r.URL.Query().Get("token")

	if raw != "" {
		return raw, nil
	}

	if r == nil {
		return "", ErrNilRequest
	}
	if h := r.Header.Get("Authorization"); len(h) > 7 && strings.EqualFold(h[0:7], "BEARER ") {
		raw = h[7:]
	}
	if raw == "" {
		return "", ErrTokenNotFound
	}

	return raw, nil
}

func ValidateToken(tokenStr string) (error){
	token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
        return nil, nil
    })
	claims, _ := token.Claims.(jwt.MapClaims)
	if tokenExp, ok := claims["exp"]; ok {
		if exp, ok := tokenExp.(float64); ok {
			now := time.Now().Unix()
			if int64(exp) > now {
				return nil
			}
		}
		return errors.New("cannot parse token exp")
	}
	return errors.New("token is expired")
	
}