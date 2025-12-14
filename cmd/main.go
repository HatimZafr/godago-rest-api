package main

import (
	"fmt"
	"log"
	"os"

	"godago-rest-api/docs"
	"godago-rest-api/internal/config"
	"godago-rest-api/internal/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Godago REST API
// @version 1.0.0
// @description A simple REST API built with Gin, MySQL, and OpenAPI
// @host localhost:8080
// @BasePath /

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL must be set in .env file")
	}

	host := os.Getenv("HOST")
	if host == "" {
		host = "127.0.0.1"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	bindAddress := fmt.Sprintf("%s:%s", host, port)

	log.Println("Connecting to database...")
	dbConfig, err := config.NewDatabaseConfig(databaseURL)
	if err != nil {
		log.Fatalf("Failed to create database connection pool: %v", err)
	}
	defer dbConfig.Close()

	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	router.Use(cors.New(corsConfig))

	routes.SetupRoutes(router, dbConfig.GetDB())

	docs.SwaggerInfo.Title = "Godago REST API"
	docs.SwaggerInfo.Description = "A simple REST API built with Gin, MySQL, and OpenAPI"
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.Host = bindAddress
	docs.SwaggerInfo.BasePath = "/"

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Printf("Starting server at http://%s", bindAddress)
	log.Printf("Swagger UI available at http://%s/swagger/index.html", bindAddress)

	if err := router.Run(bindAddress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
