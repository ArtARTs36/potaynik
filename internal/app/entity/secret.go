package entity

import (
	"fmt"
	"time"
)

type Secret struct {
	Key         string                `json:"key"`
	Value       string                `json:"value"`
	TTL         int                   `json:"ttl"`
	AuthFactors map[string]AuthFactor `json:"auth_factors"`
}

func (s *Secret) HasAuthFactors() bool {
	return len(s.AuthFactors) > 0
}

func (s *Secret) Duration() time.Duration {
	dur, _ := time.ParseDuration(fmt.Sprintf("%ss", s.TTL))

	return dur
}

func (s *Secret) AuthFactorNames() []string {
	names := make([]string, 0, len(s.AuthFactors))

	for name := range s.AuthFactors {
		names = append(names, name)
	}

	return names
}
