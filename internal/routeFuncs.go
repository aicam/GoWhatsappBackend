package internal

import (
	"github.com/aicam/secure-messenger/internal/cryptoUtils"
	"github.com/gin-gonic/gin"
	"github.com/thanhpk/randstr"
	"log"
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
		res.Info.PublicKey = s.ServerKey.PublicKeyStr
		srcUserToken := s.getUserToken(body.Info.SrcUsername)
		destUserToken := s.getUserToken(body.Info.DestUsername)
		if srcUserToken == "" || body.Info.TokenRequested {
			var newToken string
			if destUserToken != "" {
				s.setUserToken(body.Info.SrcUsername, destUserToken)
				newToken = destUserToken
			} else {
				newToken = randstr.String(16)
				s.setUserToken(body.Info.SrcUsername, newToken)
				s.setUserToken(body.Info.DestUsername, newToken)
			}
			clientPublicKey, err := cryptoUtils.ParseRsaPublicKeyFromPemStr(body.Info.PublicKey)
			if err != nil {
				res.ErrorText = "Wrong PublicKey"
				c.JSON(400, res)
				return
			}
			res.Info.SessionKey = (cryptoUtils.EncryptWithPublicKey(newToken, clientPublicKey))
			log.Print(newToken)
			c.JSON(http.StatusOK, res)
			return
		}
		// message
		if body.Message.Text != "" {
			mess, err := cryptoUtils.DecryptAES([]byte(srcUserToken), body.Message.Text)
			if err == nil {
				err = s.addMessage(Messages{
					SrcUsername:  body.Info.SrcUsername,
					DestUsername: body.Info.DestUsername,
					Text:         mess,
				})
				if err != nil {
					log.Print(err)
				}
			} else {
				log.Print(err)
			}
		}
		if body.Message.Offset == 0 {

		}
		// /message
	}
}
