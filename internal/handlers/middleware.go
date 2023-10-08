package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func (h Handlers) ValidateUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		headers := c.Request().Header

		value, ok := headers["Authorization"]
		if !ok {
			return c.JSON(http.StatusUnauthorized, requestResponse{Message: "no authorization header"})
		}

		headerParts := strings.Split(value[0], " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, requestResponse{Message: "invalid authorization header"})
		}

		if len(headerParts[1]) == 0 {
			return c.JSON(http.StatusUnauthorized, requestResponse{Message: "empty token"})
		}

		_, err := jwt.Parse(headerParts[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(h.secret), nil
		})

		if err != nil {
			return c.JSON(http.StatusUnauthorized, requestResponse{Message: "invalid token"})
		}

		return next(c)
	}
}
