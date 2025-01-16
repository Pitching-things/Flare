package middleware

import (
	"time"

	"github.com/Pitching-things/Flare/helper"
	"github.com/gin-gonic/gin"
)

func Middleware(c *gin.Context) {
	var flag int
	cookie, err := c.Cookie("user")

	if err == nil {
		Claims, er := helper.DataOfJwt(cookie)
		if er != nil {
			flag = 1
		} else {
			c.Set("user", Claims.Id)
		}

	}

	if err != nil || flag == 1 {
		userId := helper.IdCreate(5)
		jwtKey, er := helper.JwtCreate(userId)

		if er != nil {
			c.JSON(500, "error while creating jwt.")
		}

		c.SetCookie("user", jwtKey, int((time.Hour * 24).Seconds()), "/", "localhost", false, true)
		c.Set("user", userId)
	}

}
