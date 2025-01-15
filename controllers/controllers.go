package controllers

import (
	"encoding/json"
	"time"

	"github.com/Pitching-things/Flare/config"
	"github.com/Pitching-things/Flare/helper"
	"github.com/Pitching-things/Flare/models"
	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func RoomCreate(c *gin.Context) {
	RoomId := helper.IdCreate(10)
	UserId := c.MustGet("user").(string)

	Data := models.Rooms{
		CreatedAt: time.Now(),
		RoomId:    RoomId,
		HostId:    UserId,
	}

	if err := config.Rb.Set(RoomId, Data, 24*time.Hour); err.Err() != nil {
		c.JSON(500, "error while storing in redis")
	}

	if err := helper.QrCreator("http://localhost:8080/join?id="+RoomId, RoomId); err != nil {
		c.JSON(500, "error while creating room qr")
	}

	c.Redirect(308, "http://localhost:8080/room/"+RoomId)
}

func RoomJoin(c *gin.Context) {
	RoomId := c.Query("room")
	UserId := c.MustGet("user").(string)

	var Data models.Rooms

	data, err := config.Rb.Get(RoomId).Result()
	if err != nil {
		c.JSON(500, "error while fetching from redis")
	}

	if err := json.Unmarshal([]byte(data), &Data); err != nil {
		c.JSON(500, "error while decoding redis jsom data")
	}

	Data.Members = append(Data.Members, UserId)

	if err := config.Rb.Set(RoomId, Data, 24*time.Hour); err.Err() != nil {
		c.JSON(500, "error while updating members in room on redis")
	}

	c.Redirect(308, "http://localhost:8080/room/"+RoomId)
}
