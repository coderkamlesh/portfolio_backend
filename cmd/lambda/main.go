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
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambdaV2

func init() {
	gin.SetMode(gin.ReleaseMode)

	cfg := config.LoadConfig()
	config.InitDB(cfg)
	config.DB.AutoMigrate(&model.User{}, &model.Project{}, &model.Skill{}, &model.Experience{})

	r := gin.New()
	r.Use(gin.Recovery())

	// CORS enable karo taaki headers properly pass ho sakein (Authorization, Content-Type, etc.)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Production mein specific domain dalna
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
	}))

	// Debug middleware - yeh check karne ke liye ki request aa rahi hai ya nahi
	r.Use(func(c *gin.Context) {
		log.Printf("[GIN] %s %s | Headers: %v", c.Request.Method, c.Request.URL.Path, c.Request.Header)
		c.Next()
	})

	// Health check
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Portfolio API running on Lambda",
			"status":  "ok",
		})
	})

	// Agar tumhari routes /api prefix ke saath register hain toh yahan use karo
	// otherwise routes.SetupRoutes(r) directly chalega
	routes.SetupRoutes(r)

	ginLambda = ginadapter.NewV2(r)
	log.Println("Lambda initialization complete")
}

func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	// Debug logging
	log.Printf("[LAMBDA] RawPath: %s | Path: %s | Method: %s", req.RawPath, req.RequestContext.HTTP.Path, req.RequestContext.HTTP.Method)
	log.Printf("[LAMBDA] Headers: %v", req.Headers)

	// Important: Agar API Gateway stage name path mein add kar raha hai (e.g., /prod/api/...)
	// toh usse remove karna padega. HTTP API usually stage name nahi add karta $default stage pe.

	response, err := ginLambda.ProxyWithContext(ctx, req)

	// Ensure headers properly return ho rahe hain
	if response.Headers == nil {
		response.Headers = make(map[string]string)
	}
	response.Headers["Content-Type"] = "application/json"

	return response, err
}

func main() {
	lambda.Start(Handler)
}
