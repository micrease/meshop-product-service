package repository

import "github.com/jinzhu/gorm"

type Base struct {
	db *gorm.DB
}
