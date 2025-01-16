package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Rooms struct {
	CreatedAt time.Time `json:"createdat"`
	RoomId    string    `json:"roomid"`
	HostId    string    `json:"hostid"`
	Members   []string  `json:"users"`
}

type Claims struct {
	Id interface{}
	jwt.StandardClaims
}
