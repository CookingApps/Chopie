package main

import (
	"Chopie/internal/config"
	"Chopie/internal/model"
	"Chopie/internal/repository"
	"fmt"
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
	fmt.Print(ut)

	if ut != nil {
		panic(ut)
	}
	Router := gin.Default()
	Router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	Router.Run(":" + cfg.Port)
}
