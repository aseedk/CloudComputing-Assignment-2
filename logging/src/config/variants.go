package config

import (
	"github.com/joho/godotenv"
	"os"
)

var (
	_        = godotenv.Load()
	MongoURI = os.Getenv("MONGO_URI")
)
