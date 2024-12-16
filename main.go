package main

import (
	"github.com/Pitching-things/Flare/config"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.SetUp(r)
}
