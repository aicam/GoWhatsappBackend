package internal

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) test() gin.HandlerFunc {
	return func(c *gin.Context) {
		_ = s.RedisClient.Set("a", "1234", 1000000000)
		arg, _ := s.RedisClient.Get("a").Result()

		c.JSON(200, arg)
	}
}


