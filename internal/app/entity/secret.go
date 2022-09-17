package entity

type Secret struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	TTL   int    `json:"ttl"`
}
