package models

type Address struct {
	HouseNo int
	City    string
	State   string
}

type User struct {
	ID      string `json:"Id"` 
	Name    string `json:"Name"`
	Age     int
	Contact int
	Email   string
	Address Address
	Role    string `json:"role"; enum:"normal_user, admin_user"` // role is used for telling the type of participation in field
}
