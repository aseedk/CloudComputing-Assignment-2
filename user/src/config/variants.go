package config

import (
	"github.com/joho/godotenv"
	"os"
)

var (
	_               = godotenv.Load()
	MongoURI        = os.Getenv("MONGO_URI")
	LoggingURI      = os.Getenv("LOGGING_SERVICE_URL")
	OrganizationURI = os.Getenv("ORGANIZATION_SERVICE_URL")
)
