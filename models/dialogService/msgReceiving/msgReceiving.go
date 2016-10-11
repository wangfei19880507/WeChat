package msgReceiving

// Verification verifies validity of message.
type Verification struct {
	Signature string `json:"signature"`
	TimeStamp string `json:"time_stamp"`
	Nonce     string `json:"nonce"`
	EchoStr   string `json:"echostr"`
}
