package config

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func SetUp(r *gin.Engine) {

	r.Use(cors.Default())
	LoadEnv()
	DbConn()

	r.LoadHTMLGlob("templates/*")

}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to env......")
	}
}

func DbConn() {
	db, err := gorm.Open(postgres.Open(os.Getenv("PSQL_URL")))

	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Connect to Psql")

	db.AutoMigrate()

	Db = db
}
