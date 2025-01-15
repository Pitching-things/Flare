package config

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var Rb *redis.Client

func SetUp(r *gin.Engine) {

	r.Use(cors.Default())
	LoadEnv()

	Rb = RedisConn()

	r.LoadHTMLGlob("assets/templates/*")
}

func LoadEnv() {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Failed to env......")
	}
}

func RedisConn() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("REDIS_ADR"),
		Password: viper.GetString("REDIS_PASS"),
		DB:       viper.GetInt("REDIS_DB"),
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		log.Fatal("Could not connect to Redis: ", err)
	} else {
		fmt.Println("Connected to Redis")
	}

	return rdb
}
