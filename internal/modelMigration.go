package internal

import (
	"log"
	"time"
	"github.com/jinzhu/gorm"
)

type Messages struct {
	SrcUsername string `json:"src_username"`
	DestUsername string `json:"dest_username"`
}


func MakeMigrations(connectionString string) *gorm.DB {
	var db *gorm.DB
	var err error
	retryCount := 10
	for {
		db, err = gorm.Open("mysql", connectionString)
		if err != nil {
			log.Print("mysql connection error : ", err)
			time.Sleep(2 * time.Second)
		} else {
			break
		}
		if retryCount == 0 {
			log.Print("maximum retry reached")
			break
		}
		retryCount--
	}
	db.AutoMigrate(&MessagesIdentity{})
	db.AutoMigrate(&Messages{})
	db.AutoMigrate(&KeyTable{})
	return db
}
