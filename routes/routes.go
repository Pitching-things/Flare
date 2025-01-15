package routes

import (
	"github.com/Pitching-things/Flare/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {

	// Homes
	r.GET("/", controllers.HomePage)

}
