package api

import "time"

type Model struct {
	ID        uint64    `gorm:"column:id;primary_key;auto_increment;" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;not null;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;not null;" json:"updated_at"`
}

type Customer struct {
	Model
	CustomerID  string `gorm:"column:customer_id;" json:"customer_id"`
	Order       string `gorm:"column:order;" json:"order"`
	UserName    string `gorm:"column:user_name;" json:"user_name"`
	Code        string `gorm:"column:code;" json:"code"`
	Email       string `gorm:"column:email;" json:"email"`
	Password    string `gorm:"column:password;" json:"password"`
	PhoneNumber string `gorm:"column:phone_number;" json:"phone_number"`
}
