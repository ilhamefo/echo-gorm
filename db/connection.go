package db

import (
	"gorm-echo/env"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (res *gorm.DB) {
	env.LoadEnv()

	hostname := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSL_MODE")
	timezone := os.Getenv("DB_TIMEZONE")

	dsn := "host=" + hostname + " user=" + username + " password=" + password + " dbname=" + db_name + " port=" + port + " sslmode=" + sslmode + " TimeZone=" + timezone
	res, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database: " + err.Error())
		return
	}

	return res
}
