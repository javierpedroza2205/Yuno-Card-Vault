package middleware

import (
	"net/http"
	microErrors "yuno-cards/api-gateway/utils"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type Authorization struct {}

func (s *Authorization) Authorize() echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			request := context.Request()

			tokenRaw, err := FromHeaderQuery(request)

			if err != nil {
				errorResponse := microErrors.Error{
					Status: http.StatusText(400),
					Error: microErrors.ErrorReason{
						Type:        "Validate token",
						StatusCode:  401,
						Message:     err.Error(),
						UserMessage: "Authentication token is invalid, malformed, expired, has an invalid signature or an invalid algorithm.",
						Code:        "auth_bad_request",
					},
				}
				return context.JSON(http.StatusUnauthorized, errorResponse)
			}
			errt := ValidateToken(tokenRaw)

			if errt != nil {
				errorResponse := microErrors.Error{
					Status: http.StatusText(400),
					Error: microErrors.ErrorReason{
						Type:        "Validate token",
						StatusCode:  401,
						Message:     "expired token",
						UserMessage: "Authentication token is invalid, malformed, expired, has an invalid signature or an invalid algorithm.",
						Code:        "auth_bad_request",
					},
				}
				return context.JSON(http.StatusUnauthorized, errorResponse)
			}

			return next(context)
		}
	}
}

func GetUserToken(context echo.Context, key string) (string){
	request := context.Request()

	tokenRaw, _ := FromHeaderQuery(request)

	token, _ := jwt.Parse(tokenRaw, func(token *jwt.Token) (interface{}, error) {
        return nil, nil
    })
	claims, _ := token.Claims.(jwt.MapClaims)
	tokenUser := claims[key].(string)

	return tokenUser
}