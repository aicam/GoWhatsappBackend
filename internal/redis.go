package internal

import (
	"github.com/go-redis/redis/v7"
	"time"
)





var ExpirationTime = 1200  * 1000000000

func GetClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "@Ali@021021", // no password set
		DB:       0,  // use default DB
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

func (s *Server) getMessage(srcUsername string, destUsername string) string {
	val, err := s.RedisClient.Get(srcUsername + destUsername + "message").Result()
	if err == redis.Nil {
		return ""
	}
	s.RedisClient.Del(srcUsername + destUsername + "message")
	return val
}

func (s *Server) setUserToken(username string, token string)  {
	s.RedisClient.Set(username + "Token", token, time.Duration(ExpirationTime))
}

func (s *Server) setUsersMessage(srcUsername string, destUsername string, message string) string {
	s.RedisClient.Set(srcUsername + destUsername + "message", message, 0)
}