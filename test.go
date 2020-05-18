package main

import (
	"fmt"
	"github.com/aicam/secure-messenger/internal/cryptoUtils"
	"github.com/go-redis/redis/v7"
	"log"
	"time"
)

type JsStruct struct {
	Hash []byte `json:"hash"`
}

func ExampleClient() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "@Ali@021021", // no password set
		DB:       0,             // use default DB
	})
	err := client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	err = client.Set("keyv", "value", 1000000000).Err()
	if err != nil {
		panic(err)
	}
	time.Sleep(10 * time.Second)
	err = client.Set("keye", "value", 1000000000).Err()
	if err != nil {
		panic(err)
	}
	// Output: key value
	// key2 does not exist
}
func main() {
	//log.Print(randstr.Base64(25))
	p, _ := cryptoUtils.GenerateKeyPair(256)
	log.Print(cryptoUtils.ExportRsaPrivateKeyAsPemStr(p))
	ExampleClient()
}
