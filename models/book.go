package models

import (
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Book struct {
	ID     uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Title  string    `json:"title"`
	Author string    `json:"author"`
}

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateBookInput struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
}

func (book Book) BeforeCreate(scope *gorm.DB) error {
	scope.Statement.SetColumn("id", uuid.NewV4().String())
	return nil
}
