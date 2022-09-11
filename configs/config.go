package configs

import (
	"assignment-2/models"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitConfig() models.Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config := models.Config{}
	config.WebServer.Port = os.Getenv("PORT")

	config.Database.Host = os.Getenv("PGSQL_HOST")
	config.Database.Name = os.Getenv("PGSQL_NAME")
	config.Database.Password = os.Getenv("PGSQL_PASSWORD")
	config.Database.Port = os.Getenv("PGSQL_PORT")
	config.Database.User = os.Getenv("PGSQL_USER")

	return config
}
