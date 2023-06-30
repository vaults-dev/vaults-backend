package initializers

import (
	"os"

	"github.com/vaults-dev/vaults-backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlConnectDB() *gorm.DB {
	gormDB, err := gorm.Open(mysql.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return gormDB
}

func MigrateTable(gormDB *gorm.DB) {
	gormDB.AutoMigrate(&models.User{})
}
