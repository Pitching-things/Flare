package models

import "time"

type Rooms struct {
	CreatedAt time.Time `json:"createdat"`
	RoomId    string    `json:"roomid"`
	HostId    string    `json:"hostid"`
	Members   []string  `json:"users"`
}
