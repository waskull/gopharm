package domain

import "time"

type User struct {
	Model
	Fullname  string     `json:"fullname" binding:"required"`
	Email     string     `json:"email" gorm:"unique" binding:"required"`
	Password  string     `json:"password" binding:"required"`
	BirthDate *time.Time `json:"created_at" gorm:"default:current_timestamp"`
	Dni       string     `json:"dni" binding:"required"`
	Phone     string     `json:"phone" binding:"optional"`
	Role      string     `json:"role" binding:"required"`
}
