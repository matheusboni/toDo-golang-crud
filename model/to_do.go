package model

import "time"

type ToDo struct {
	Id          string `gorm:"primary_key" json:"id"`
	Title       string `gorm:"size:255;not null" json:"title" binding:"required"`
	Description string `gorm:"size:255;not null" json:"description" binding:"required"`
	Status      string `gorm:"size:255;not null" json:"status" binding:"required"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
}