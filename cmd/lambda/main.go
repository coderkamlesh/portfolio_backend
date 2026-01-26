package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/coderkamlesh/portfolio_backend/config"
	"github.com/coderkamlesh/portfolio_backend/internal/http/routes"
	"github.com/coderkamlesh/portfolio_backend/internal/model"
	"github.com/gin-gonic/gin"
)

// Global variable for the adapter
var ginLambda *ginadapter.GinLambda

// init() function Lambda start hone par ek baar chalta hai (Cold Start optimization)
func init() {
	log.Println("Gin server initializing...")

	// 1. Config Load
	cfg := config.LoadConfig()

	// 2. Db Connect
	config.InitDB(cfg)

	// 3. Migrate Tables (Lambda me ye har cold start pe check karega)
	config.DB.AutoMigrate(&model.User{}, &model.Project{}, &model.Skill{}, &model.Experience{})

	// 4. Gin Setup
	r := gin.Default()

	// Health Check Route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "Server running on AWS Lambda",
			"mode":   gin.Mode(),
		})
	})

	// 5. Route Setup (Tumhare existing routes)
	routes.SetupRoutes(r)

	// 6. Wrap Gin router in Lambda Adapter
	ginLambda = ginadapter.New(r)
}

// Handler function receives the event from API Gateway and passes it to Gin
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Agar ginLambda initialize nahi hua (rare case), toh init karo
	if ginLambda == nil {
		log.Fatal("GinLambda adapter not initialized")
	}

	// Request ko Gin router pe proxy karo
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	// Lambda execution start
	lambda.Start(Handler)
}
