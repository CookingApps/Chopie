package main

import (
	"Chopie/internal/config"
	"Chopie/internal/handler"
	"Chopie/internal/model"
	"Chopie/internal/repository"
	"Chopie/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)

	}

	db := repository.InitDB(cfg)
	ut := db.AutoMigrate(&model.User{})

	authRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(authRepo)
	authHandle := handler.NewAuthHandler(*authService)

	if ut != nil {
		panic(ut)

	}
	Router := gin.Default()
	Router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	Router.POST("/Register", authHandle.CreateUser)
	Router.Run(":" + cfg.Port)
}
