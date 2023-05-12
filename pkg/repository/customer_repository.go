package repository

import (
	"fmt"

	"github.com/AbdulrahmanDaud10/slade360-customer-order-service/pkg/api"
	"gorm.io/gorm"
)

type Storage interface {
	CreateOrder(customer api.Customer) error
}

type storage struct {
	db *gorm.DB
}

func NewStorage(db *gorm.DB) Storage {
	return &storage{
		db: db,
	}
}

func (s *storage) CreateCustomerService(customer api.Customer) error {
	return nil
}

func (s *storage) CreateOrder(customer api.Customer) error {
	if customer := s.db.Create(&customer); customer.Error != nil {
		return fmt.Errorf(customer.Error.Error())
	}
	return nil
}
