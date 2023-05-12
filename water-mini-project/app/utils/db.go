package utils

import (
	"water-mini-project/app/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
func InitializeDB() error {
    dsn := "313147:lyoko1234567@tcp(mysql-water-mini.alwaysdata.net)/water-mini_project?charset=utf8mb4&parseTime=True&loc=Local"
    conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return err
    }
    db = conn	

	db.AutoMigrate(&models.Customer{})
	db.AutoMigrate(&models.Bill{})
    
    return nil
}
func GetDB() *gorm.DB {
    return db
}