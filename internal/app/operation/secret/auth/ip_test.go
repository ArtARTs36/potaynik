package auth

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewIP(t *testing.T) {
	cases := []struct {
		address       string
		expectedType  int
		expectedError error
	}{
		{
			address:      "172.22.0.1",
			expectedType: ipTypeAddress,
		},
		{
			address:      "127.0.0.0/8",
			expectedType: ipTypeNetwork,
		},
	}

	for _, tCase := range cases {
		ip, err := newIP(tCase.address)

		assert.Equal(t, tCase.expectedError, err)

		if err == nil {
			assert.Equal(t, tCase.expectedType, ip.Type)
		}
	}
}
