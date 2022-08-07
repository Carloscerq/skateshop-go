package models

import (
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
    "SkateShop/utils"
    "os"
    "strconv"
)

var dbConnection *gorm.DB

func InitDB() {
    var err error
    dbURL, ok := os.LookupEnv("DATABASE_URL"); if !ok {
        panic("DATABASE_URL not found")
    }

    dbConnection, err = gorm.Open(mysql.Open(dbURL), &gorm.Config{}); if err != nil {
        panic("failed to connect database")
    }

    // Handle migrations
    dbConnection.AutoMigrate(&User{})

    sqlDB, err := dbConnection.DB()
    defer sqlDB.Close()
    if err != nil {
        panic("failed to handle database")
    }

    maxIdlleConns, err := strconv.Atoi(utils.GetEnv("MAX_IDLE_CONNS", "10")); if err != nil {
        panic("failed to read MAX_IDLE_CONNS")
    }
    maxOpenConns, err := strconv.Atoi(utils.GetEnv("MAX_OPEN_CONNS", "10")); if err != nil {
        panic("failed to read MAX_OPEN_CONNS")
    }

    sqlDB.SetMaxIdleConns(maxIdlleConns)
    sqlDB.SetMaxOpenConns(maxOpenConns)
}
