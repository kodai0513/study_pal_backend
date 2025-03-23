package middlewares

import (
	"net/http"
	"strings"
	"study-pal-backend/app/app_types"
	"study-pal-backend/app/utils/study_pal_jwts"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AuthRequired(devModeLogin bool, secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if devModeLogin {
			c.Set("user_id", uuid.MustParse("1b142a47-765f-46ce-be5c-5d37c8ffbca5"))
		} else {
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

			userId, err := study_pal_jwts.VerifyToken(secretKey, tokenString)

			if err != nil {
				c.SecureJSON(
					http.StatusUnauthorized,
					app_types.NewErrorResponse([]string{err.Error()}),
				)
				c.Abort()
				return
			}

			c.Set("user_id", userId)
		}
		c.Next()
	}
}
