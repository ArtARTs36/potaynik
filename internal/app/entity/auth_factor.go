package entity

type AuthFactor struct {
	Key    string                 `json:"key"`
	Params map[string]interface{} `json:"params"`
}
