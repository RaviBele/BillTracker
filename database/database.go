package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect() error {
	var err error

	Database, err = gorm.Open(
		mysql.Open(os.Getenv("DNS")),
	)

	if err == nil {
		fmt.Println("Successfully connected to MySQL!")
	}

	return err
}
