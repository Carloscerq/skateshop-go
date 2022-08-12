package models

import (
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
    "os"
)

var DbConnection *gorm.DB

func InitDB() {
    var err error
    dbURL, ok := os.LookupEnv("DATABASE_URL"); if !ok {
        panic("DATABASE_URL not found")
    }

    DbConnection, err = gorm.Open(mysql.Open(dbURL), &gorm.Config{}); if err != nil {
        panic("failed to connect database")
    }

    // Handle migrations
    DbConnection.AutoMigrate(&User{})
    DbConnection.AutoMigrate(&Product{})
}
