package infrastructure

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase(dbUser, dbPassword, dbHost, dbPort, dbName string) (*gorm.DB, error) {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
