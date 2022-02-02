package jwt

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("JWTAuth")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error no JWTAuth cookie": err,
			})
			return
		}

		jid, username, err := validateToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error cookie JWTAuth invalid": err,
			})
			return
		}
		c.Set("jid", jid)
		c.Set("username", username)

		c.Next()
	}
}
