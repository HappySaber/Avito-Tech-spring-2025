package midlleware

import (
	"PVZ/src/utils"

	"github.com/gin-gonic/gin"
)

func IsAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("token")

		if err != nil {
			c.JSON(401, gin.H{"error": "Couldn't get cookie 'token'"})
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(cookie)

		if err != nil {
			c.JSON(401, gin.H{"error": "Couldn't parse the token: " + err.Error()})
			c.Abort()
			return
		}
		c.Set("role", claims.Role)
		c.Next()
	}
}

func IsModerator() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role != "moderator" {
			c.JSON(403, gin.H{"error": "Access denied, only moderator can do this"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func IsPVZemployee() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role != "PVZemployee" {
			c.JSON(403, gin.H{"error": "Access denied, only PVZemployee can do this"})
			c.Abort()
			return
		}
		c.Next()
	}
}
