package handlers

import (
	"errors"
	"fmt"
	"net/http"

	jwt "github.com/golang-jwt/jwt/v5"
)

func Authentication(h CustomHandler) CustomHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		token := r.Header.Get("Authentication")
		if token == "" {
			return errors.New("unauthorized")
		}

		return h(w, r)
	}
}

const jwtSecret = "aksjdflaksdjf"

func parseJwtToken(tokenStr string) error {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return jwtSecret, nil
	})

	if !token.Valid {
		switch {
		case errors.Is(err, jwt.ErrTokenSignatureInvalid):
			return errors.New("invalid token signature")
		case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
			return errors.New("token not valid yet")
		default:
			return errors.New("that's not a valid token")
		}
	}

	claims, ok := token.Claims.(jwt.MapClaim)
	if !ok {
		return errors.New("token invalid")
	}

	_ = claims["token"]

	return err
}
