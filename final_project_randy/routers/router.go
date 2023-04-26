package routers

import (
	"project-my-gram/controllers"
	_ "project-my-gram/docs"
	"project-my-gram/middlewares"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Mygram API
// @version 1.0
// @description Pada final project ini, kalian akan diminta untuk membuat suatu aplikasi bernama MyGram, yang dimana pada aplikasi ini kalian dapat menyimpan foto maupun membuat comment untuk foto orang lain.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

func StartApp(c controllers.UserRepo, p controllers.PhotoRepo, o controllers.CommentRepo, m controllers.MediaRepo) *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", c.UserRegister)
		userRouter.POST("/login", c.UserLogin)
	}

	photoRouter := r.Group("/photo")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.GET("/getphotos", p.GetPhoto)
		photoRouter.GET("/getphotobyid/:photoId", p.GetPhotoById)
		photoRouter.POST("/uploadphoto", p.UploadPhoto)
		photoRouter.PUT("/updatephoto/:photoId", middlewares.PhotoAuthorization(), p.UpdatePhoto)
		photoRouter.DELETE("/deletephoto/:photoId", middlewares.PhotoAuthorization(), p.DeletePhoto)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.GET("/getcomments", o.GetComment)
		commentRouter.GET("/getcommentbyid/:commentId", o.GetCommentById)
		commentRouter.POST("/uploadcomment", o.UploadComment)
		commentRouter.PUT("/updatecomment/:commentId", middlewares.CommentAuthorization(), o.UpdateComment)
		commentRouter.DELETE("/deletecomment/:commentId", middlewares.CommentAuthorization(), o.DeleteComment)
	}

	mediaRouter := r.Group("/socialmedia")
	{
		mediaRouter.Use(middlewares.Authentication())
		mediaRouter.GET("/getsocmeds", m.GetMedia)
		mediaRouter.GET("/getsocmedbyid/:socialMediaId", m.GetSocialMediabyId)
		mediaRouter.POST("/uploadsocmed", m.UploadMedia)
		mediaRouter.PUT("/updatesocmed/:socialMediaId", middlewares.MediaAuthorization(), m.UpdateMedia)
		mediaRouter.DELETE("deletesocmed/:socialMediaId", middlewares.MediaAuthorization(), m.DeleteMedia)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}
