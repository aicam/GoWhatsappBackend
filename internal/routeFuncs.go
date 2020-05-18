package internal

import (
	"github.com/aicam/secure-messenger/internal/cryptoUtils"
	"github.com/gin-gonic/gin"
	"github.com/thanhpk/randstr"
	"net/http"
)

func (s *Server) test() gin.HandlerFunc {
	return func(c *gin.Context) {
		body := SendRequestStruct{}
		res := SendResponseStruct{}
		err := c.BindJSON(&body)
		if err != nil || body.Info.PublicKey == "" {
			c.JSON(400, struct {
				Error string `json:"error"`
			}{Error: "Body format does not supported"})
		}
		srcUserToken := s.getUserToken(body.Info.SrcUsername)
		destUserToken := s.getUserToken(body.Info.DestUsername)
		if srcUserToken == "" {
			newToken := randstr.String(9)
			if destUserToken != "" {
				//s.setUserToken(body.Info.SrcUsername, destUserToken)
			} else {
				//newToken := randstr.String(54)
				//s.setUserToken(body.Info.SrcUsername, newToken)
				//s.setUserToken(body.Info.DestUsername, newToken)
			}
			res.Info.PublicKey = s.ServerKey.PublicKeyStr
			clientPublicKey, err := cryptoUtils.ParseRsaPublicKeyFromPemStr(body.Info.PublicKey)
			if err != nil {
				c.JSON(400, err)
			}
			res.Info.SessionKey = string(cryptoUtils.EncryptWithPublicKey([]byte(newToken), clientPublicKey))
			c.JSON(http.StatusOK, res)
		}
	}
}
