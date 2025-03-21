package main

import (
	"bytes"
	"cloud-computing/users/config"
	"cloud-computing/users/database"
	"cloud-computing/users/restful/models/dao"
	"cloud-computing/users/restful/route"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"time"
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

	// Apply middleware globally
	r.Use(LoggingMiddleware())

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})
	r.GET("/api/", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})
	route.SetupUserRoute(r)

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

type responseRecorder struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r *responseRecorder) Write(b []byte) (int, error) {
	if r.body == nil {
		r.body = bytes.NewBufferString("")
	}
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

// LoggingMiddleware sends request details to the logging service
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// Capture request body
		var requestBodyBytes []byte
		if c.Request.Body != nil {
			requestBodyBytes, _ = io.ReadAll(c.Request.Body)
		}
		c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBodyBytes)) // Restore body for next handlers

		// Process the request
		c.Next()

		// Capture response body
		writer := c.Writer
		respRecorder := responseRecorder{ResponseWriter: writer, body: bytes.NewBuffer(make([]byte, 0, 16*1024))}
		c.Writer = &respRecorder

		// Prepare log data
		logData := map[string]interface{}{
			"url":            c.Request.URL.Path,
			"method":         c.Request.Method,
			"userId":         "default", // Extract from headers or session
			"organizationId": "default", // Extract from headers
			"requestBody":    string(requestBodyBytes),
			"requestQuery":   c.Request.URL.RawQuery,
			"responseBody":   respRecorder.body.String(),
			"timestamp":      startTime.Format(time.RFC3339),
		}

		// Send log to logging service
		go func() {

			if logData["requestBody"] == "" || logData["responseBody"] == "" {
				return
			}
			loggingServiceURL := config.LoggingURI // Update with actual URL

			jsonData, err := json.Marshal(logData)
			if err != nil {
				log.Println("Error marshalling log data:", err)
			}
			_, err = http.Post(loggingServiceURL, "application/json", bytes.NewBuffer(jsonData))
			if err != nil {
				log.Println("Error sending log to logging service:", err)
			}
		}()
	}
}
