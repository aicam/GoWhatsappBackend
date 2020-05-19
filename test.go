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
	//js, _ := json.Marshal(internal.SendRequestStruct{
	//	Info: struct {
	//		TokenRequested bool   `json:"token_requested"`
	//		SrcUsername    string `json:"src_username"`
	//		DestUsername   string `json:"dest_username"`
	//		PublicKey      string `json:"public_key"`
	//	}{
	//		TokenRequested: true,
	//		SrcUsername: "aicam",
	//		DestUsername: "sijal",
	//		PublicKey: "-----BEGIN PUBLIC KEY-----\nMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDJaAuS55Ix1Cx04Ysc1k4SDQ6k++YHcV3wiZNZIykx4e4mjPTY2VdXTAxlJefJsNSJydD+7hmto3zRf+kwiSBp7UFGYMMfjqTx31kypjBkId2mxr2gD7PwUmF1WgzHjgpzICGxPporZlqnhOePeEA8afgi2+By1j583bGz/NmKEQIDAQAB\n-----END PUBLIC KEY-----",
	//	},
	//	Message: struct {
	//		Text   string `json:"text"`
	//		Offset int    `json:"offset"`
	//		HMAC   string `json:"hmac"`
	//	}{
	//		Text: "Hello world!",
	//	},
	//	File: struct {
	//		Data     string `json:"data"`
	//		Chunk    string `json:"chunk"`
	//		Finished bool   `json:"finished"`
	//		FileName string `json:"file_name"`
	//		HMAC     string `json:"hmac"`
	//	}{},
	//})
	//log.Print(string(js))
	log.Print(cryptoUtils.EncryptAES([]byte("pvchJ2OVCO8YlHI9"), "hello world"))
	sv, _ := cryptoUtils.DecryptCBC([]byte("pvchJ2OVCO8YlHI9"), []byte("89GQWAMj4QOcyjhmEuijUw=="))
	log.Print(string(sv))
}
