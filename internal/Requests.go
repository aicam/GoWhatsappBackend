package internal

type SendRequestStruct struct {
	Info struct {
		TokenRequested bool   `json:"token_requested"`
		SrcUsername    string `json:"src_username"`
		DestUsername   string `json:"dest_username"`
		PublicKey      string `json:"public_key"`
		RequestedFile  string `json:"requested_file"`
		RequestedChunk int    `json:"requested_chunk"`
	} `json:"info"`
	Message struct {
		Text   string `json:"text"`
		Offset int    `json:"offset"`
		HMAC   string `json:"hmac"`
	} `json:"message"`
	File struct {
		Data     string `json:"data"`
		Chunk    int    `json:"chunk"`
		Finished bool   `json:"finished"`
		FileName string `json:"file_name"`
		HMAC     string `json:"hmac"`
	} `json:"file"`
}
