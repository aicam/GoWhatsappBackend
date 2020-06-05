package internal

import (
	"github.com/aicam/secure-messenger/internal/cryptoUtils"
	"github.com/gin-gonic/gin"
)

func (s *Server) checkAuthentication(f func()) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		username := c.GetHeader("username")
		pass := s.getUserPassword(username)

	}
}
