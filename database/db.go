package database

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Konek() {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASS"),
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_NAME"),
    )

    var err error
    DB, err = sql.Open("mysql", dsn)
    if err != nil {
        log.Fatalf("error konek: %v", err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatalf("ga dijangkau: %v", err)
    }

    log.Println("konek")
}
