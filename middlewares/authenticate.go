package middlewares

import (
	"api-jwt/helpers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtData, err := helpers.ValidateToken(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
				"message": err.Error(),
			})
			return
		}

		c.Set("userData", jwtData)
		fmt.Println(jwtData)
		c.Next()
	}
}