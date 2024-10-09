package repo

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"urlshortner/config"
	"urlshortner/internal/models"
)

func InitializeDB() *gorm.DB {
	host := config.GetDBConfig("db.host")
	user := config.GetDBConfig("db.user")
	password := config.GetDBConfig("db.password")
	dbName := config.GetDBConfig("db.name")
	port := config.GetDBConfig("db.port")
	sslmode := config.GetDBConfig("db.sslmode")
	timezone := config.GetDBConfig("db.timezone")

	customDbUrl := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, password, dbName, port, sslmode, timezone,
	)

	db, err := gorm.Open(postgres.Open(customDbUrl), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	err = db.AutoMigrate(&models.URL{})
	if err != nil {
		return nil
	}

	return db
}
