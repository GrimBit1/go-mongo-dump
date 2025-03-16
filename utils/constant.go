package utils

import (
	"encoding/hex"
	"encoding/json"
	"net"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const uri = "mongodb://root:example@localhost:27017/"

type User struct {
	Id        bson.ObjectID `json:"_id" bson:"_id"`
	FirstName string        `json:"first_name" bson:"first_name"`
	LastName  string        `json:"last_name" bson:"last_name"`
	Email     string        `json:"email" bson:"email"`
	Gender    string        `json:"gender" bson:"gender"`
	IpAddress net.IP        `json:"ip_address" bson:"ip_address"`
}

// NormalizeIP converts IpAddress if it was stored as a hex string
func (u *User) NormalizeIP() net.IP {
	ipStr := u.IpAddress.String()
	// Trim '?' if present.
	if len(ipStr) > 0 && ipStr[0] == '?' {
		ipStr = ipStr[1:]
	}
	// Try parsing normally.
	ip := net.ParseIP(ipStr)
	// Fallback to decoding hex if the standard parse fails.
	if ip == nil {
		if decoded, err := hex.DecodeString(ipStr); err == nil {
			ip = net.ParseIP(string(decoded))
		}
	}
	u.IpAddress = ip
	return u.IpAddress
}

func (u User) MarshalJSON() ([]byte, error) {
	// Create a copy and normalize its IP.
	user := u
	(&user).NormalizeIP()

	// Alias to prevent recursion.
	type Alias User
	return json.Marshal(&struct {
		IpAddress string `json:"ip_address"`
		*Alias
	}{
		IpAddress: user.IpAddress.String(),
		Alias:     (*Alias)(&user),
	})
}
