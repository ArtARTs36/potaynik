package entity

type Secret struct {
	Key         string                `json:"key"`
	Value       string                `json:"value"`
	TTL         int                   `json:"ttl"`
	AuthFactors map[string]AuthFactor `json:"auth_factors"`
}

func (s *Secret) HasAuthFactors() bool {
	return len(s.AuthFactors) > 0
}
