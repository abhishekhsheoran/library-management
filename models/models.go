package models

import (
	"time"
)

type Address struct {
	HouseNo int
	City    string
	State   string
}

type User struct {
	ID        string
	Name      string
	Age       int
	Contact   int
	Email     string
	Password  string
	Address   Address
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Book struct {
	Name      string
	Auther    string
	IssuedBy  []string
	NumOfCopy int
}
