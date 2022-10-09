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

type CreateTodoInput struct {
	Text   string `json:"text"`
	Done   bool   `json:"done"`
	UserID uint   `json:"userId"`
}

type UpdateTodoInput struct {
	Text   *string `json:"text"`
	Done   *bool   `json:"done"`
	UserID *uint   `json:"userId"`
}

type Todo struct {
	gorm.Model
	Text   string `json:"text"`
	Done   bool   `json:"done"`
	UserID uint   `json:"userId"`
	User   User   `json:"user"`
}
