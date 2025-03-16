package utils

import "net"

const uri = "mongodb://root:example@localhost:27017/"

type User struct {
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	Email     string `json:"email" bson:"email"`
	Gender    string `json:"gender" bson:"gender"`
	IpAddress net.IP `json:"ip_address" bson:"ip_address"`
}
