package internal

type SendRequestStruct struct {
	Info struct {
		SrcUsername  string `json:"src_username"`
		DestUsername string `json:"dest_username"`
		PublicKey    string `json:"public_key"`
	} `json:"info"`
	Message struct {
		Text   string `json:"text"`
		Offset int    `json:"offset"`
		HMAC   string `json:"hmac"`
	} `json:"message"`
	File struct {
		Data     string `json:"data"`
		Chunk    string `json:"chunk"`
		Finished bool   `json:"finished"`
		FileName string `json:"file_name"`
		HMAC     string `json:"hmac"`
	} `json:"file"`
}
