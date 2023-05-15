package api

import (
	"errors"
	"regexp"
	"strconv"

	"go.starlark.net/lib/time"
)

type orderCustomerRequest struct {
	Item       string    `josn:"item"`
	Amount     string    `json:"amount"`
	Time       time.Time `json:"time"`
	CustomerID string    `json:"customer_id"`
}

type CustomerService interface {
	AddOrder(customerID string, request orderCustomerRequest) error
	validateAmount(amount string) error
}

type customerRepository interface {
	CreateOrder(customer Customer) error
}

type customerService struct {
	storage customerRepository
}

func NewCustomerService(customerRepo customerRepository) CustomerService {
	return &customerService{storage: customerRepo}
}

func (c *customerService) AddOrder(customerID string, request orderCustomerRequest) error {
	if err := c.validateAmount(request.Amount); err != nil {
		return err
	}

	newOrder := Customer{
		Model:       Model{},
		CustomerID:  customerID,
		Order:       "",
		Code:        "",
		PhoneNumber: "",
	}

	if err := c.storage.CreateOrder(newOrder); err != nil {
		return err
	}
	return nil
}

func (c *customerService) validateAmount(amount string) error {
	if amount == "" {
		return errors.New("amount cannot be blank")
	}

	value := regexp.MustCompile("[+]?([0-9]*[.])?[0-9]+")
	if !value.MatchString(amount) {
		return errors.New("value must be a number")
	}

	newamount, _ := strconv.ParseFloat(amount, 64)
	if newamount < 0 {
		return errors.New("cannot be a negative")
	}
	return nil
}
