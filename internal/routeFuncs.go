package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhpk/randstr"
)

func (s *Server) test() gin.HandlerFunc {
	return func(c *gin.Context) {
		body := SendRequestStruct{}
		res := SendResponseStruct{}
		err := c.BindJSON(&body)
		if err != nil {
			c.JSON(400, struct {
				Error string `json:"error"`
			}{Error: "Body format does not supported"})
		}
		userToken := s.getUserToken(body.Info.SrcUsername)
		if userToken == "" {
			newToken := randstr.String(54)
			s.setUserToken(body.Info.SrcUsername, newToken)
			s.setUserToken(body.Info.DestUsername, newToken)
			res.Info.PublicKey = s.ServerKey.PublicKeyStr
			res.Info.SessionKey =
		}
	}
}


