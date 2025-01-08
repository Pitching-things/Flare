package main

import (
	"os"

	"github.com/Pitching-things/Flare/config"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.SetUp(r)

	r.Run(os.Getenv("PORT"))
}
