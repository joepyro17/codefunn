package databases

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	. "pyro.com/codefunn/models"
)

var DB *gorm.DB

func DbConnection(){
	var err error

	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	DbHostname := os.Getenv("POSTGRES_HOST")
	DbUser := os.Getenv("POSTGRES_USER")
	DbPassword := os.Getenv("POSTGRES_PASSWORD")
	DbName := os.Getenv("POSTGRES_DB")
	DbPort := os.Getenv("POSTGRES_PORT")

	dsn := "host="+ DbHostname +" user=" + DbUser + " password=" + DbPassword + " dbname=" + DbName + " port=" + DbPort + " sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Migrate schemas
	err = DB.AutoMigrate(&Blog{}, &Category{})
	if err != nil {
		log.Fatal("Failed in db.AutoMigrate")
	}

	fmt.Println("DB Connection Success")

}