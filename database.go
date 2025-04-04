package main

import (
    "database/sql"
    "log"
    "os" // Para trabajar con el sistema de archivos

    _ "github.com/mattn/go-sqlite3" // Driver SQLite
)

// nuestros modelos, estos representan las tablas de la base de datos. Los pueden poner en un archivo models.go

type UserModel struct {
    Username string `json:"username"`
    Password string `json:"password"` // Needed from the client request
}

// setupDatabase inicializa la conexión a la BD
func setupDatabase(dbPath string) (*sql.DB, error) {
    log.Printf("Conectando a la base de datos en: %s", dbPath)
    db, err := sql.Open("sqlite3", dbPath)
    if err != nil {
        return nil, err
    }

    // Es buena idea hacer ping para verificar la conexión inmediatamente
    if err = db.Ping(); err != nil {
        db.Close() // Cerrar si el ping falla
        return nil, err
    }

    log.Println("Base de datos conectada exitosamente.")
    // Podríamos añadir aquí la creación de tablas si no existen
    // _, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (...)`)

    return db, nil
}
