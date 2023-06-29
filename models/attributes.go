package models



type Address struct {
	HouseNo int
	City    string
	State   string
}



type user struct {
	ID      string
	Name    string
	Age     int
	Contact int
	Email   string
	Address Address
	Role    string
}

