package internal

import "github.com/jinzhu/gorm"

func addMessage(message Messages, db *gorm.DB) error {
	err := db.Save(&message).Error
	if err != nil {
		return err
	}
	return nil
}

func addFileData(fileData FilesData, db *gorm.DB) error {
	err := db.Save(&fileData).Error
	if err != nil {
		return err
	}
	return nil
}
