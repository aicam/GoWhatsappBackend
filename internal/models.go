package internal

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

func (s *Server) getMessageDB(srcUsername string, destUsername string, limit int) []Messages {
	var returnMessages []Messages
	s.DB.Order("id desc").Limit(limit).Where(&Messages{
		SrcUsername:  srcUsername,
		DestUsername: destUsername,
	}).Find(&returnMessages)
	return returnMessages
}
