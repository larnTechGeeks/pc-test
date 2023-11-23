package dtos

type MessageRequest struct {
	Text string `json:"text"`
}

type Message struct {
	ID   int64  `json:"id"`
	Text string `json:"text"`
	Spam string `json:"spam"`
}

type MessageResposne struct {
	Result string `json:"result"`
}
