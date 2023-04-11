package middlewares

import (
	"api-jwt/database"
	"api-jwt/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var UserRoleAuth string

func ProductAuthorizationById() gin.HandlerFunc{
	return func(c *gin.Context) {
		db := database.GetDB()
		productId ,err := strconv.Atoi(c.Param("productId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Bad Request",
				"message": "invalid parameter",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		
		Product := models.Product{}

		err = db.Select("user_id").First(&Product, uint(productId)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "Data Not Found",
				"message": "data doesn't exist",
			})
			return
		}

		if UserRoleAuth == "user" {
			if Product.UserID != userID {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": "Unauthorized",
					"message": "you are not allowed to access this data",
				})
				return
			}
		}

		c.Next()
	}
}
func ProductAuthorizationAll() gin.HandlerFunc{
	return func(c *gin.Context) {
		db := database.GetDB()
		productId ,err := strconv.Atoi(c.Param("productId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Bad Request",
				"message": "invalid parameter",
			})
			return
		}

		if UserRoleAuth != "admin" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
				"message": "you are not allowed to access this data",
			})
			return
		}
		
		Product := models.Product{}

		err = db.Select("user_id").First(&Product, uint(productId)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "Data Not Found",
				"message": "data doesn't exist",
			})
			return
		}
		c.Next()
	}
}

func UserRoleAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		userData := c.MustGet("userData").(jwt.MapClaims)
		User := models.User{}
		err := db.First(&User, "id = ?", userData["id"]).Error
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Not Found",
				"message": err.Error(),
			})
			return
		}
		UserRoleAuth = User.Role
		c.Next()
	}
}