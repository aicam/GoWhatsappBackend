package internal

import (
	"github.com/aicam/secure-messenger/internal/cryptoUtils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) checkAuthentication(f gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		username := c.GetHeader("username")
		pass := s.getUserPassword(username)
		if cryptoUtils.CheckJWTToken(token, []byte(pass)) != nil {
			c.JSON(http.StatusUnauthorized, struct {
				Status string `json:"status"`
			}{Status: "Unauthorized"})
		} else {
			f(c)
		}

	}
}
