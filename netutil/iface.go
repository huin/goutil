package netutil

import (
	"bytes"
	"errors"
	"net"
)

func IPForAddr(addr net.Addr) (net.IP, error) {
	switch addr := addr.(type) {
	case *net.IPNet:
		return addr.IP, nil
	case *net.IPAddr:
		return addr.IP, nil
	case *net.TCPAddr:
		return addr.IP, nil
	case *net.UDPAddr:
		return addr.IP, nil
	default:
		return net.IP{}, errors.New("unknown address type")
	}
}

func IsZeroIP(ip net.IP) bool {
	switch len(ip) {
	case 0:
		return true
	case net.IPv4len:
		return bytes.Equal([]byte(ip), net.IPv4zero)
	case net.IPv6len:
		return bytes.Equal([]byte(ip), net.IPv6zero)
	}
	return false
}

func ExpandIPAddr(ip net.IP) ([]net.IP, error) {
	if !IsZeroIP(ip) {
		return []net.IP{ip}, nil
	}
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	ips := make([]net.IP, 0, len(addrs))
	for _, addr := range addrs {
		ifaceIP, err := IPForAddr(addr)
		if err != nil {
			// Ignore address - does not have an IP address.
			continue
		}
		ips = append(ips, ifaceIP)
	}
	return ips, nil
}
