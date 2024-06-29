package database

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
    "os"
)

var DB *gorm.DB

func ConnectDB() {
    dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
        os.Exit(1)
    }
}
