package models

import "time"

type User struct {
	//gorm.Model
	ID             string     `json:"id,readonly",gorm:"primaryKey,column:_id"`
	DocumentTypeID string     `json:"documentTypeId",gorm:"foreignKey:Users"`
	Document       string     `json:"document",gorm:"column:document"`
	Name           string     `json:"name",gorm:"column:name"`
	LastName       string     `json:"lastName",gorm:"column:last_name"`
	Email          string     `json:"email",gorm:"column:email" validate:"required,email`
	CreatedAt      time.Time  `json:"createdAt,readonly"`
	UpdatedAt      time.Time  `json:"updatedAt,readonly"`
	DeletedAt      *time.Time `json:"deletedAt,readonly",sql:"index"`
}
