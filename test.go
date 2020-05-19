package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/aicam/secure-messenger/internal/cryptoUtils"
	"log"
)

func Decrypt(key []byte, ct string) {
	ciphertext, _ := hex.DecodeString(ct)
	c, err := aes.NewCipher(key)
	if err != nil {
		_ = fmt.Errorf("NewCipher(%d bytes) = %s", len(key), err)
		panic(err)
	}
	plain := make([]byte, len(ciphertext))
	c.Decrypt(plain, ciphertext)
	s := string(plain[:])
	fmt.Printf("AES Decrypyed Text:  %s\n", s)
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
	//log.Print(cryptoUtils.EncryptAES([]byte("pvchJ2OVCO8YlHI9"), "hello world"))
	//sv, _ := cryptoUtils.DecryptCBC([]byte("pvchJ2OVCO8YlHI9"), []byte("89GQWAMj4QOcyjhmEuijUw=="))
	//log.Print(string(sv))

	ciphertext, err := base64.StdEncoding.DecodeString("lVVRybGxbp8xXj23g5x7RC3JZ24s2VXN6TBnQ922b6c=")
	if err != nil {
		log.Fatal(err)
	}
	key := []byte("pvchJ2OVCO8YlHI9")
	log.Print(cryptoUtils.DecryptAES(key, "A5242EFC97A87A6926AB4688603409FD"))
	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}
	gcmDecrypt, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}
	nonceSize := gcmDecrypt.NonceSize()
	if len(ciphertext) < nonceSize {
		fmt.Println(err)
	}
	nonce, encryptedMessage := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcmDecrypt.Open(nil, nonce, encryptedMessage, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(plaintext))
}
