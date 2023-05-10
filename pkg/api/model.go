package api

import "time"

type Model struct {
	ID        uint64    `gorm:"column:id;primary_key;auto_increment;" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;not null;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;not null;" json:"updated_at"`
}

type Customer struct {
	Model
	customerID  string `gorm:"column:customer_id;" json:"customer_id"`
	order       string `gorm:"column:order;" json:"order"`
	userName    string `gorm:"column:user_name;" json:"user_name"`
	code        string `gorm:"column:code;" json:"code"`
	email       string `gorm:"column:email;" json:"email"`
	password    string `gorm:"column:password;" json:"password"`
	phoneNumber string `gorm:"column:phone_number;" json:"phone_number"`
}
