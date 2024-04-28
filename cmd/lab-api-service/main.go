package main

import (
	"log"
	"os"
	"strings"

	"github.com/Magum23/wac-lab-webapi/api"
	"github.com/Magum23/wac-lab-webapi/internal/lab"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Printf("Server started")
	port := os.Getenv("LAB_API_PORT")
	if port == "" {
		port = "8080"
	}
	environment := os.Getenv("LAB_API_ENVIRONMENT")
	if !strings.EqualFold(environment, "production") { // case insensitive comparison
		gin.SetMode(gin.DebugMode)
	}
	engine := gin.New()
	engine.Use(gin.Recovery())
	// request routings
	lab.AddRoutes(engine)
	engine.GET("/openapi", api.HandleOpenApi)
	engine.Run(":" + port)
}
