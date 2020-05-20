package internal

import (
	"github.com/go-redis/redis/v7"
	"time"
)

var ExpirationTime = 60 * 60 * 1000000000
var MessageExpiration = 20 * time.Minute

func GetClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "@Ali@021021", // no password set
		DB:       0,             // use default DB
	})
	err := client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
	return client
}

func (s *Server) getUserToken(username string) string {
	val, err := s.RedisClient.Get(username + "Token").Result()
	if err == redis.Nil {
		return ""
	}
	return val
}

func (s *Server) getMessageRedis(srcUsername string, destUsername string) []Messages {
	returnMessages := []Messages{}
	for i := 0; i < 100; i++ {
		val, err := s.RedisClient.Get(srcUsername + destUsername + "message" + string(i)).Result()
		if err == redis.Nil {
			break
		}
		returnMessages = append(returnMessages, Messages{
			SrcUsername:  srcUsername,
			DestUsername: destUsername,
			Text:         val,
		})
		s.RedisClient.Del(srcUsername + destUsername + "message" + string(i))
	}
	return returnMessages
}

func (s *Server) setUserToken(username string, token string) {
	s.RedisClient.Set(username+"Token", token, time.Duration(ExpirationTime))
}

func (s *Server) setUsersMessageRedis(srcUsername string, destUsername string, message string) {
	for i := 0; i < 100; i++ {
		_, err := s.RedisClient.Get(srcUsername + destUsername + "message" + string(i)).Result()
		if err == redis.Nil {
			s.RedisClient.Set(srcUsername+destUsername+"message", message+string(i), MessageExpiration)
			break
		}
	}
}
