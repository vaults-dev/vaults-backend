package initializers

import (
	"github.com/vaults-dev/vaults-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBconn *gorm.DB

func ConnectDB() {
	var err error
	dsn := "postgres://postgres:postgrespw@localhost:55000"
	DBconn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	MigrateTable()
}

func MigrateTable() {
	DBconn.AutoMigrate(&models.User{})
}
