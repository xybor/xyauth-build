package models

import (
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	UserID string
	Secret string
}
