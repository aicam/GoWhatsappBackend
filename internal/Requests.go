package internal

type RequestStruct struct {
	Info struct{
		SrcUsername string `json:"src_username"`
		DestUsername string `json:"dest_username"`
		PublicKey string `json:"public_key"`
	} `json:"info"`
	Message struct{
		Text string `json:"text"`
		Offset int `json:"offset"`
	} `json:"message"`
	File struct{
		Data string `json:"data"`
		Chunk string `json:"chunk"`
		Finished bool `json:"finished"`
	} `json:"file"`
}
