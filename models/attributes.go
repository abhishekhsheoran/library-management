package models

type Address struct {
	HouseNo int
	City    string
	State   string
}

type User struct {
	ID      string  
	Name    string 
	Age     int
	Contact int
	Email   string
	Address Address
	Role    string
}
