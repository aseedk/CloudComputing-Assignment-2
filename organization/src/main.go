package main

import (
	"cloud-computing/organization/organization/src/config"
	"cloud-computing/organization/organization/src/database"
	"cloud-computing/organization/organization/src/restful/models/dao"
	"cloud-computing/organization/organization/src/restful/route"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	// Load environment variables
	mongoURI := config.MongoURI
	if mongoURI == "" {
		log.Fatal("MONGO_URI environment variable is required")
	}

	// Connect to MongoDB
	if err := database.ConnectMongo(mongoURI); err != nil {
		log.Fatalf("Could not connect to MongoDB: %v", err)
	}
	defer database.CloseMongo()

	err := dao.InitMongoDB()
	if err != nil {
		return
	}
	// Initialize Gin router
	r := gin.Default()

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})
	route.SetupOrganizationRoute(r)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}
	log.Printf("Server running on port %s", port)
	err = r.Run(":" + port)
	if err != nil {
		log.Println("Error starting server:", err)
		return
	}
}
