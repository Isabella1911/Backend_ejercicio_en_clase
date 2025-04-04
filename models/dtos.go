            package models

// --- Request DTOs ---

// LoginRequest defines the expected structure for the login request body.
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// RegisterRequest defines the expected structure for the registration request body.
// For this simple case, it's identical to LoginRequest.
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// --- Response Data DTOs ---

// LoginSuccessData defines the structure for the data returned on successful login.
type LoginSuccessData struct {
	UserID   int64  `json:"userId"`
	Username string `json:"username"`
}

// RegisterSuccessData defines the structure for the data returned on successful registration.
type RegisterSuccessData struct {
	UserID   int64  `json:"userId"`
	Username string `json:"username"`
}

// --- Generic API Response Structures ---

// ErrorDetail provides a structured way to return error information.
type ErrorDetail struct {
	Message string `json:"message"`
	// Code    string `json:"code,omitempty"`    // Optional: Application-specific error code
	// Field   string `json:"field,omitempty"`   // Optional: Field related to a validation error
}

// APIResponse is a generic wrapper for standardizing API responses.
// It clearly indicates success and includes either data or error details.
type APIResponse struct {
	Success bool         `json:"success"`
	Data    any          `json:"data,omitempty"`    // Use 'any' (Go 1.18+) or interface{} for flexibility
	Error   *ErrorDetail `json:"error,omitempty"` // Pointer allows omitting the field if there's no error
}

// --- Factory Functions for API Responses ---

// NewSuccessResponse creates a standard success response with data.
// It acts like a constructor for successful APIResponse instances.
func NewSuccessResponse(data any) APIResponse {
	return APIResponse{
		Success: true,
		Data:    data,
		Error:   nil, // Explicitly nil, will be omitted by json tag
	}
}

// NewErrorResponse creates a standard error response.
// It acts like a constructor for failed APIResponse instances.
func NewErrorResponse(errorMessage string) APIResponse {
	return APIResponse{
		Success: false,
		Data:    nil, // Explicitly nil, will be omitted
		Error: &ErrorDetail{
			Message: errorMessage,
		},
	}
}
