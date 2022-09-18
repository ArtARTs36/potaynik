package responses

import (
	"encoding/json"
)

type Response struct {
	Code    int
	Message []byte
	Headers map[string]string
}

type errorMessage struct {
	Error string `json:"error"`
}

func newErrorResponseMessage(message string) []byte {
	jsonMsg, _ := json.Marshal(errorMessage{Error: message})

	return jsonMsg
}
