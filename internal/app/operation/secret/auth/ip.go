package auth

import "net"

const (
	ipTypeAddress = iota
	ipTypeNetwork
)

type IP struct {
	Value string `json:"value"`
	Type  int    `json:"type"`
}

func newIP(address string) (*IP, error) {
	netIp, _, err := net.ParseCIDR(address)

	if err == nil {
		return &IP{
			Value: netIp.String(),
			Type:  ipTypeNetwork,
		}, nil
	}

	// todo нужна валидация
	netIp = net.ParseIP(address)

	return &IP{
		Value: netIp.String(),
		Type:  ipTypeAddress,
	}, nil
}

func (ip *IP) Contains(addr string) bool {
	thatNetIP := net.ParseIP(addr)

	if ip.Type == ipTypeAddress {
		selfNetIP := net.ParseIP(ip.Value)

		return selfNetIP.Equal(thatNetIP)
	}

	_, network, _ := net.ParseCIDR(ip.Value)

	return network.Contains(thatNetIP)
}
