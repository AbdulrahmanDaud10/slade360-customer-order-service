package repository

import "gorm.io/gorm"

type Storage interface {
}

type storage struct {
	db *gorm.DB
}

func NewStorage(db *gorm.DB) Storage {
	return &storage{
		db: db,
	}
}
