package db

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
    "log"
    "os"
)

var DB *sql.DB

func Conectar() {
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")

    connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    var err error
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal("Error al conectar a la base de datos:", err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatal("No se pudo hacer ping a la BD:", err)
    }

    log.Println(" Conexión a la base de datos exitosa")
}
