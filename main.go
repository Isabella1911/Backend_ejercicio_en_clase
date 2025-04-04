package main

import (
    "log"
    "net/http"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

func main() {
    // Conectar a la base de datos
    db, err := setupDatabase("./users.db")
    if err != nil {
        log.Fatal("CRITICAL: No se pudo conectar a la base de datos:", err)
    }
    defer db.Close() // Asegurar que se cierre al final

    // Crear router Chi
    r := chi.NewRouter()

    // Middlewares
    r.Use(middleware.Logger)    // Loggea cada request
    r.Use(middleware.Recoverer) // Recupera de panics
    r.Use(configureCORS())      // Aplica nuestra configuración CORS

    // Rutas Públicas (sin autenticación requerida inicialmente)
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("API de Login v1.0"))
    })
    r.Post("/register", postRegisterHandler(db))
    r.Post("/login", postLoginHandler(db))

    // Ruta para obtener datos de usuario (protegida más tarde con JWT)
    // Por ahora, cualquiera puede acceder si conoce el ID
    r.Get("/users/{userID}", getUserHandler(db))


    // Iniciar servidor
    port := ":3000"
    log.Printf("Servidor escuchando en puerto %s", port)
    log.Fatal(http.ListenAndServe(port, r))
}
