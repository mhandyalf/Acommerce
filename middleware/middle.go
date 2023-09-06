package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func JWTAuth(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {
		tokenHeader := c.Request().Header.Get("Authorization")
		if tokenHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Token tidak ada"})
		}

		tokenParts := strings.Split(tokenHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Format token salah"})
		}

		tokenString := tokenParts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Gantilah "secret-key" dengan kunci rahasia yang sesuai yang digunakan saat membuat token
			return []byte("secret-key"), nil
		})

		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Token tidak valid"})
		}

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			c.Set("username", claims["username"]) // Mengirimkan klaim username ke konteks
			return next(c)
		}

		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Token tidak valid"})
	}
}
