package database

import (
    "myapp/app"
)

func Migrate() {
    DB.AutoMigrate(&app.User{})
    DB.AutoMigrate(&app.Photo{})
}
