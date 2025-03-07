package middlewares

import (
	"net/http"
	"strings"
	"study-pal-backend/app/app_types"
	"study-pal-backend/app/utils/study_pal_jwts"

	"github.com/gin-gonic/gin"
)

func AuthRequired(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.SecureJSON(
				http.StatusUnauthorized,
				app_types.NewErrorResponse([]string{"authorization header missing"}),
			)
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		isValid, _, err := study_pal_jwts.VerifyToken(secretKey, tokenString)

		if err != nil || !isValid {
			c.SecureJSON(
				http.StatusUnauthorized,
				app_types.NewErrorResponse([]string{"invalid token"}),
			)
			c.Abort()
			return
		}

		c.Next()
	}
}
