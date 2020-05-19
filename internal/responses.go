package internal

import "time"

type SendResponseStruct struct {
	Status    bool   `json:"status"`
	ErrorText string `json:"error_text"`
	Info      struct {
		PublicKey  string `json:"public_key"`
		SessionKey []byte `json:"session_key"`
	} `json:"info"`
	ReturnMessages []struct {
		Text   string    `json:"text"`
		Date   time.Time `json:"date"`
		Sender string    `json:"sender"`
		HMAC   string    `json:"hmac"`
	} `json:"return_messages"`
	ReturnFileData struct {
		Data     string `json:"data"`
		Chunk    int    `json:"chunk"`
		Key      string `json:"key"`
		FileName string `json:"file_name"`
		HMAC     string `json:"hmac"`
	}
}
