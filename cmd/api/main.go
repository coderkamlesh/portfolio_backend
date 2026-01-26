package main

import (
	"fmt"
	"log"

	"github.com/coderkamlesh/portfolio_backend/config"
	"github.com/coderkamlesh/portfolio_backend/internal/http/routes"
	"github.com/coderkamlesh/portfolio_backend/internal/model"
	"github.com/gin-gonic/gin"
)

func main() {
	//  Config Load karo sabse pehle
	cfg := config.LoadConfig()
	//  Db
	config.InitDB(cfg)

	//tables migrate
	config.DB.AutoMigrate(&model.User{}, &model.Project{}, &model.Skill{}, &model.Experience{})

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "Server running....",
			"mode":   gin.Mode(),
		})
	})
	//route setup
	routes.SetupRoutes(r)

	serverAddr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Starting server on port %s in %s mode...", cfg.Port, cfg.GinMode)

	if err := r.Run(serverAddr); err != nil {
		log.Fatal("Server start failed: ", err)
	}
}
