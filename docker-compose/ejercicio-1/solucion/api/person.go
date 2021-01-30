package main

type Person struct {
	ID          int    `bson:"_id,omitempty" json:"_id,omitempty"`
	FirstName   string `bson:"first_name,omitempty" json:"first_name,omitempty"`
	LastName    string `bson:"last_name,omitempty" json:"last_name,omitempty"`
	Company     string `bson:"company,omitempty" json:"company,omitempty"`
	Email       string `bson:"email,omitempty" json:"email,omitempty"`
	IPAddress   string `bson:"ip_address,omitempty" json:"ip_address,omitempty"`
	PhoneNumber string `bson:"phone,omitempty" json:"phone,omitempty"`
}
