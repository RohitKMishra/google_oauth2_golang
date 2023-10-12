package initializers

import (
	"fmt"
	"log"

	"github.com/RohitKMishra/google_oauth2_golang/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("golang.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}

	DB.AutoMigrate(&models.User{})
	fmt.Println("Connected Successfully to the Database")
}
