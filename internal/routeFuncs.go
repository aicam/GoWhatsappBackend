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
			res.Status = false
			res.ErrorText = err.Error()
			c.JSON(400, res)
			return
		}
		res.Info.PublicKey = s.ServerKey.PublicKeyStr
		srcUserToken := s.getUserToken(body.Info.SrcUsername)
		destUserToken := s.getUserToken(body.Info.DestUsername)
		// New token
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
		if body.Message.Offset != -1 {
			res.ReturnMessages = s.getMessageDB(body.Info.SrcUsername, body.Info.DestUsername, body.Message.Offset, srcUserToken)
		} else {
			res.ReturnMessages = s.getMessageRedis(body.Info.SrcUsername, body.Info.DestUsername)
		}
		res.Status = true
		log.Print(srcUserToken)
		c.JSON(http.StatusOK, res)
		// /message
		if body.File.FileName != "" {
			if cryptoUtils.ComputeHmac256(body.File.Data[:100], srcUserToken) != body.File.HMAC {
				return
			}
			s.DB.Save(FilesData{
				SrcUsername:  body.Info.SrcUsername,
				DestUsername: body.Info.DestUsername,
				FileName:     body.File.FileName,
				Data:         body.File.Data,
				Chunk:        body.File.Chunk,
				Key:          srcUserToken,
			})
		}
		if body.Message.Text != "" {
			if cryptoUtils.ComputeHmac256(body.Message.Text, srcUserToken) != body.Message.HMAC {
				log.Print(cryptoUtils.ComputeHmac256(body.Message.Text, srcUserToken))
				return
			}
			s.setUsersMessageRedis(body.Info.SrcUsername, body.Info.DestUsername, body.Message.Text)
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
	}
}
