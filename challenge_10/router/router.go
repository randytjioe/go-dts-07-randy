package router

import (
	"github.com/gin-gonic/gin"
	"github.com/randytjioe/go-dts-07-randy/challenge_10/controllers"
	"github.com/randytjioe/go-dts-07-randy/challenge_10/middlewares"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	productRouter := r.Group("/products")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/", controllers.CreateProduct)
		productRouter.GET("/getallproduct", controllers.GetAllProduct)
		productRouter.PUT("/:productId", middlewares.ProductAuthorization(), controllers.UpdateProduct)
		productRouter.GET("/getproduct/:productId", controllers.GetProductById)
		productRouter.DELETE("/deleteproduct/:productId", middlewares.ProductAuthorization(), controllers.DeleteProduct)
	}

	return r
}
