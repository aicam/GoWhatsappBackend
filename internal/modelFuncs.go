package internal

import "github.com/aicam/secure-messenger/internal/cryptoUtils"

func (s *Server) addMessage(message Messages) error {
	err := s.DB.Save(&message).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) addFileData(fileData FilesData) error {
	err := s.DB.Save(&fileData).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) getMessageDB(srcUsername string, destUsername string, limit int, key string) []Messages {
	var returnMessages []Messages
	s.DB.Order("id desc").Limit(limit).Where("src_username IN (?) AND dest_username IN (?) ",
		[]string{srcUsername, destUsername}, []string{srcUsername, destUsername}).Find(&returnMessages)
	for i, item := range returnMessages {
		returnMessages[i].Text = cryptoUtils.EncryptAES([]byte(key), item.Text)
	}
	return returnMessages
}
