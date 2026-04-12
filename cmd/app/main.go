package main

import (
	"Chopie/internal/config"
	"Chopie/internal/repository"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg, err := config.LoadConfig()

	db := repository.InitDB(cfg)
	fmt.Print(db, "Omo")

	if err != nil {
		panic(err)
	}
	Router := gin.Default()
	Router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	Router.Run(":" + cfg.Port)
}
