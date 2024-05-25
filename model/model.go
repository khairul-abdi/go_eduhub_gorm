package model

import (
	"time"

	"gorm.io/gorm"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type Student struct {
	ID        uint      `gorm:"primarykey"`
	FullName  string    `gorm:"full_name"`
	Address   string    `json:"address"`
	Class     string    `json:"class"`
	IsActive  int       `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedBy string    `json:"updated_by"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedBy string    `json:"deleted_by"`
	DeletedAt time.Time `json:"deleted_at"`
}

type Course struct {
	gorm.Model
	Name       string `json:"name"`
	Schedule   string `json:"schedule"`
	Attendance int    `json:"attendance"`
	IsActive   int    `json:"is_active"`
	DeletedBy  string `json:"deleted_by"`
	UpdatedBy  string `json:"updated_by"`
	CreatedBy  string `json:"created_by"`
}

type Credential struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
	Port         int
	Schema       string
}
