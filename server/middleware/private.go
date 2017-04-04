package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/javinc/mango/server/auth"
)

// PrivateMiddleware checks if has valid token
func PrivateMiddleware(checkPayload func(map[string]interface{}) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		// validates token
		p, err := auth.CheckToken(
			getToken(c.Request.Header.Get("authorization")))

		// validates payload
		if err == nil {
			err = checkPayload(p)
		}

		if err != nil {
			c.String(400, err.Error())
			c.Abort()

			return
		}

		c.Next()
	}
}

func getToken(authHeader string) string {
	s := strings.Split(strings.TrimSpace(authHeader), " ")
	if len(s) != 2 {
		return ""
	}

	return s[1]
}
