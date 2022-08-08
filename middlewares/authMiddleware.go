package middlewares

import (
	"mitramas_test/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Request.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			// c.Abort()
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		valueToken := tokenString.Value
		if valueToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "request does not contain an access token"})
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		errs := auth.ValidateToken(valueToken)
		if errs != nil {
			c.JSON(401, gin.H{"error": errs.Error()})
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		c.Next()
	}
}
