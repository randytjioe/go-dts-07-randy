package main

import (
	"project-my-gram/config"
	"project-my-gram/controllers"
	"project-my-gram/routers"
)

func main() {
	db := config.InitDB()

	userRepo := controllers.UserRepo{DB: db}
	photoRepo := controllers.PhotoRepo{DB: db}
	commentRepo := controllers.CommentRepo{DB: db}
	mediaRepo := controllers.MediaRepo{DB: db}

	r := routers.StartApp(userRepo, photoRepo, commentRepo, mediaRepo)
	r.Run(":8080")
}
