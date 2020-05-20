package main

import (
	"github.com/aicam/secure-messenger/internal"
	"log"
	"net/http"
)

const DatabaseConnectionString = "aicam:021021ali@tcp(mysql-db:3306)/messenger?charset=utf8mb4&parseTime=True"

func main() {
	s := internal.NewServer()
	db := internal.MakeMigrations(DatabaseConnectionString)
	s.DB = db
	s.RedisClient = internal.GetClient()
	s.Route()
	err := http.ListenAndServe("0.0.0.0:4300", s.Router)
	if err != nil {
		log.Print(err)
	}
}
