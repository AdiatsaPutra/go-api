package models

import (
	"time"
)

type Post struct {
	ID        uint      `json:"id" gorm:"column:id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Like      bool      `json:"like"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

// custom tablename
func (e *Post) TableName() string {
	return "m_post"
}
