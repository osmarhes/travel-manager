package authmiddleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/osmarhes/travel-manager/internal/auth"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "missing token"})
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		token, _ := jwt.ParseWithClaims(tokenStr, &auth.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if claims, ok := token.Claims.(*auth.Claims); ok && token.Valid {
			c.Set("userID", claims.UserID)
			return next(c)
		}

		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid token"})
	}
}
