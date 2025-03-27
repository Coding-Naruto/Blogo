package models

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

type Post struct {
    ID      uint   `gorm:"primaryKey"`
    Title   string `gorm:"not null"`
    Content string `gorm:"not null"`
}

var DB *gorm.DB

func InitDB() {
    var err error
    DB, err = gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect to database")
    }

    DB.AutoMigrate(&Post{})
}