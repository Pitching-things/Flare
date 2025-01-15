package main

import (
	"github.com/Pitching-things/Flare/config"
	"github.com/Pitching-things/Flare/routes"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	r := gin.Default()

	config.SetUp(r)
	routes.Routes(r)

	r.Run(viper.GetString("PORT"))
}
