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
	dns := fmt.Sprintf("%s", os.Getenv("DNS"))

	Database, err = gorm.Open(
		mysql.Open(dns),
	)

	if err == nil {
		fmt.Println("Successfully connected to MySQL!")
	}

	return err
}
