package models

import "gorm.io/gorm"

type CreateUserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateUserInput struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}
