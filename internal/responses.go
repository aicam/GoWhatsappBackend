package internal

type SendResponseStruct struct {
	Status    bool   `json:"status"`
	ErrorText string `json:"error_text"`
	Info      struct {
		PublicKey  string `json:"public_key"`
		SessionKey []byte `json:"session_key"`
	} `json:"info"`
	ReturnMessages []Messages `json:"return_messages"`
	HMAC           string     `json:"hmac"`
	ReturnFileData struct {
		Data     string `json:"data"`
		Chunk    int    `json:"chunk"`
		Key      string `json:"key"`
		FileName string `json:"file_name"`
		HMAC     string `json:"hmac"`
	}
}
