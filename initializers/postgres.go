package initializers

import (
	"os"

	"github.com/vaults-dev/vaults-backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBconn *gorm.DB

func ConnectDB() {
	var err error
	dsn := os.Getenv("DSN")
	DBconn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	MigrateTable()
}

func MigrateTable() {
	DBconn.AutoMigrate(&models.User{})
}
