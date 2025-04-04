package main

import (
    "net/http"

    "github.com/go-chi/cors"
)

// configureCORS retorna un middleware de CORS configurado
func configureCORS() func(http.Handler) http.Handler {
    // Configuración básica - ¡Ajustar en producción!
    corsMiddleware := cors.New(cors.Options{
        // Permitir orígenes específicos (ej. donde corre su frontend)
        // Usar "*" es inseguro para producción.
        AllowedOrigins:   []string{"*", "http://localhost:5500", "http://127.0.0.1:5500"}, // Añadir origen de Live Server si lo usan
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
        ExposedHeaders:   []string{"Link"},
        AllowCredentials: true, // Importante si usan cookies o auth headers
        MaxAge:           300, // Maximum value not ignored by any of major browsers
    })
    return corsMiddleware.Handler
}
