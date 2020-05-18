package internal

import (
	"crypto/rsa"
	"github.com/aicam/secure-messenger/internal/cryptoUtils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
	"log"
)

type Server struct {
	DB          *gorm.DB
	Router      *gin.Engine
	RedisClient *redis.Client
	ServerKey   struct {
		PrivateKey    *rsa.PrivateKey
		PublicKey     *rsa.PublicKey
		PrivateKeyStr string
		PublicKeyStr  string
	}
}

// Here we create our new server
func NewServer() *Server {
	router := gin.Default()
	// Here we opened cors for all
	router.Use(cors.Default())
	newServer := &Server{
		DB:          nil,
		Router:      router,
		RedisClient: nil,
	}
	newServer.ServerKey.PrivateKey, newServer.ServerKey.PublicKey = cryptoUtils.GenerateKeyPair(2048)
	newServer.ServerKey.PrivateKeyStr = cryptoUtils.ExportRsaPrivateKeyAsPemStr(newServer.ServerKey.PrivateKey)
	publicKeyStr, err := cryptoUtils.ExportRsaPublicKeyAsPemStr(newServer.ServerKey.PublicKey)
	if err != nil {
		log.Fatal(err)
	}
	newServer.ServerKey.PublicKeyStr = publicKeyStr
	return newServer
}
