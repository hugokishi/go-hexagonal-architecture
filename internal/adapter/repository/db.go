package repository

import "gorm.io/gorm"

type DB struct {
	cnn *gorm.DB
}

func NewDatabase(cnn *gorm.DB) *DB {
	return &DB{cnn}
}
