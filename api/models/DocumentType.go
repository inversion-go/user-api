package models

import "time"

type DocumentType struct {
	ID        string `json:"id,readonly",gorm:"primary_key"`
	Name      string `gorm:"column:name"`
	Users     []User
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
