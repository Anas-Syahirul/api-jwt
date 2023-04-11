package controllers

import (
	"api-jwt/database"
	"api-jwt/helpers"
	"api-jwt/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	userData := c.MustGet("userData").(jwt.MapClaims)

	Product := models.Product{}
	userID := uint(userData["id"].(float64))
	User := models.User{}
	errUser := db.First(&User, "id = ?", userID).Error
	if errUser != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Not Found",
			"message": errUser.Error(),
		})
	}

	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	Product.UserID = userID
	Product.User = &User

	err := db.Debug().Create(&Product).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Product)
}

func UpdateProduct(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	userData := c.MustGet("userData").(jwt.MapClaims)

	productId, _ := strconv.Atoi(c.Param("productId"))
	userID := uint(userData["id"].(float64))

	Product := models.Product{}
	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	Product.UserID = userID
	Product.ID = uint(productId)

	err := db.Model(&Product).Where("id = ?", productId).Updates(models.Product{
		Title:       Product.Title,
		Description: Product.Description,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Product)
}

func GetProductById(c *gin.Context) {
	db := database.GetDB()

	productId, _ := strconv.Atoi(c.Param("productId"))
	product := models.Product{}
	err := db.First(&product, "id = ?", productId).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Not Found",
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
	db := database.GetDB()
	productId, _ := strconv.Atoi(c.Param("productId"))
	product := models.Product{}
	err := db.Where("id = ?", productId).Delete(&product).Error
	if err != nil{
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Not Found",
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Product Deleted Successfully",
	})
}
