package api

import (
	"errors"
	"regexp"
	"strconv"

	"go.starlark.net/lib/time"
)

type orderCustomerRequest struct {
	item       string    `josn:"item"`
	amount     string    `json:"amount"`
	time       time.Time `json:"time"`
	customerID string    `json:"customer_id"`
}

type CustomerService interface {
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
	if err := c.validateAmount(request.amount); err != nil {
		return err
	}

	newOrder := Customer{
		Model:       Model{},
		customerID:  customerID,
		order:       "",
		userName:    "",
		code:        "",
		email:       "",
		password:    "",
		phoneNumber: "",
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
