package routers

import (
	"api-jwt/controllers"
	"api-jwt/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "HELLO!")
	})

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)

		userRouter.POST("/login", controllers.UserLogin)
	}

	productRouter := r.Group("/product")
	{
		productRouter.Use(middlewares.Authentication())

		productRouter.POST("/", controllers.CreateProduct)
		productRouter.PUT("/:productId", middlewares.UserRoleAuthorization(), middlewares.ProductAuthorizationAll(), controllers.UpdateProduct)
		productRouter.GET("/:productId", middlewares.UserRoleAuthorization(), middlewares.ProductAuthorizationById(), controllers.GetProductById)
		productRouter.DELETE("/:productId", middlewares.UserRoleAuthorization(), middlewares.ProductAuthorizationAll(), controllers.DeleteProduct)
	}

	return r
}
