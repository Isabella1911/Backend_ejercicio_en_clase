package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	// Import your models package (adjust path 'myapp' if your module name is different)
	"myapp/models"

	"golang.org/x/crypto/bcrypt"
)

// postLoginHandler handles user login attempts.
func PostLoginHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1. Decode Request Body into LoginRequest DTO from models package
		var req models.LoginRequest // Use the DTO from the models package
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Printf("Error decoding login request: %v", err)
			// Use the factory from the models package
			response := models.NewErrorResponse("Invalid request body")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}

		// 2. Basic Validation
		if req.Username == "" || req.Password == "" {
			// Use the factory from the models package
			response := models.NewErrorResponse("Username and password are required")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}

		// 3. Query Database for User ID and Hashed Password
		var storedHash string
		var userID int64 // Use int64 for database IDs
		err := db.QueryRowContext(r.Context(),
			"SELECT id, password_hash FROM users WHERE username = ?",
			req.Username,
		).Scan(&userID, &storedHash)

		if err != nil {
			// Use the factory from the models package
			response := models.NewErrorResponse("Invalid username or password") // Generic message
			statusCode := http.StatusUnauthorized

			if err != sql.ErrNoRows {
				log.Printf("Error querying user '%s': %v", req.Username, err)
				// Use the factory from the models package
				response = models.NewErrorResponse("Internal server error")
				statusCode = http.StatusInternalServerError
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(statusCode)
			json.NewEncoder(w).Encode(response)
			return
		}

		// 4. Compare Provided Password with Stored Hash
		err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(req.Password))
		if err != nil {
			// Use the factory from the models package
			response := models.NewErrorResponse("Invalid username or password")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
			return
		}

		// 5. Login Successful - Prepare and Send Success Response
		log.Printf("Login successful for user ID: %d (%s)", userID, req.Username)

		// Create the specific data payload using the DTO from the models package
		loginData := models.LoginSuccessData{ // Use the DTO from the models package
			UserID:   userID,
			Username: req.Username,
		}

		// Wrap the data in the standard APIResponse using the factory from the models package
		response := models.NewSuccessResponse(loginData)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}
