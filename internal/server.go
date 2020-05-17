package internal

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/go-redis/redis/v7"
)

type Server struct {
	DB               *gorm.DB
	Router           *gin.Engine
	RedisClient		 *redis.Client
}

// Here we create our new server
func NewServer() *Server {
	router := gin.Default()
	// Here we opened cors for all
	router.Use(cors.Default())
	return &Server{
		DB:     nil,
		Router: router,
		RedisClient: nil,
	}
}

