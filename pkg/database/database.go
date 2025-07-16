package database

import (
    "fmt"
    "log"
    "os"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"

    "travel-manager/internal/user"
    "travel-manager/internal/travel"
)

var DB *gorm.DB

func Connect() {
    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
        os.Getenv("DB_PORT"),
    )

    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }

    err = DB.AutoMigrate(&user.User{}, &travel.TravelRequest{})
    if err != nil {
        log.Fatalf("failed to run migrations: %v", err)
    }

    log.Println("📦 Database connected and migrated")
}
