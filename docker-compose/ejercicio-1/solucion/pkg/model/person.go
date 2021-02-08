package model

import "encoding/xml"

type People struct {
	XMLName xml.Name `xml:"people"`
	Person  []Person `xml:"person"`
}

type Person struct {
	ID          int    `bson:"_id,omitempty" json:"_id,omitempty" xml:"id"`
	FirstName   string `bson:"first_name,omitempty" json:"first_name,omitempty" xml:"first_name"`
	LastName    string `bson:"last_name,omitempty" json:"last_name,omitempty" xml:"last_name"`
	Company     string `bson:"company,omitempty" json:"company,omitempty" xml:"company"`
	Email       string `bson:"email,omitempty" json:"email,omitempty" xml:"email"`
	IPAddress   string `bson:"ip_address,omitempty" json:"ip_address,omitempty" xml:"ip_address"`
	PhoneNumber string `bson:"phone,omitempty" json:"phone,omitempty" xml:"phone_number"`
}
