package main

import (
	"github.com/aicam/GoWhatsappBackend/internal"
	"log"
	"net/http"
)

const DatabaseConnectionString = "aicam:12345@tcp(mysql-db:3306)/messenger?charset=utf8mb4&parseTime=True"

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
