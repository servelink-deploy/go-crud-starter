package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     *string   `json:"phone,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUserRequest struct {
	Name  string  `json:"name" binding:"required,min=1,max=255"`
	Email string  `json:"email" binding:"required,email"`
	Phone *string `json:"phone,omitempty" binding:"omitempty,max=50"`
}

type UpdateUserRequest struct {
	Name  *string `json:"name,omitempty" binding:"omitempty,min=1,max=255"`
	Email *string `json:"email,omitempty" binding:"omitempty,email"`
	Phone *string `json:"phone,omitempty" binding:"omitempty,max=50"`
}

type PaginatedResponse struct {
	Data       []User `json:"data"`
	Total      int    `json:"total"`
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
	TotalPages int    `json:"total_pages"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
