package paranoid

import (
	"net"
	"testing"
)

func TestIPv4(t *testing.T) {
	type task struct {
		result bool
		src    string
		target *net.IPNet
	}

	tests := []task{
		// Private Class
		{false, `9.255.255.255`, IPv4PrivateClassA},
		{true, `10.0.0.0`, IPv4PrivateClassA},
		{true, `10.0.0.0`, IPv4PrivateClassA},
		{true, `10.255.255.255`, IPv4PrivateClassA},
		{false, `11.0.0.0`, IPv4PrivateClassA},

		{false, `172.15.255.255`, IPv4PrivateClassB},
		{true, `172.16.0.0`, IPv4PrivateClassB},
		{true, `172.31.255.255`, IPv4PrivateClassB},
		{false, `172.32.0.0`, IPv4PrivateClassB},

		{false, `192.165.255.255`, IPv4PrivateClassC},
		{true, `192.168.0.0`, IPv4PrivateClassC},
		{true, `192.168.255.255`, IPv4PrivateClassC},
		{false, `192.169.0.0`, IPv4PrivateClassC},

		// Shared Addr, Loopback or Reserved
		{false, `100.63.255.255`, IPv4SharedAddr},
		{true, `100.64.0.0`, IPv4SharedAddr},
		{true, `100.127.255.255`, IPv4SharedAddr},
		{false, `100.128.0.0`, IPv4SharedAddr},

		{false, `126.255.255.255`, IPv4Loopback},
		{true, `127.0.0.0`, IPv4Loopback},
		{true, `127.255.255.255`, IPv4Loopback},
		{false, `128.0.0.0`, IPv4Loopback},

		{false, `169.253.255.255`, IPv4LinkLocal},
		{false, `169.253.0.0`, IPv4LinkLocal},
		{false, `169.253.255.255`, IPv4LinkLocal},
		{false, `169.255.0.0`, IPv4LinkLocal},

		{false, `191.255.255.255`, IPv4SpecialPurpose},
		{true, `192.0.0.0`, IPv4SpecialPurpose},
		{true, `192.0.0.255`, IPv4SpecialPurpose},
		{false, `192.0.1.0`, IPv4SpecialPurpose},

		{false, `239.255.255.255`, IPv4Reserved},
		{true, `240.0.0.0`, IPv4Reserved},
		{true, `255.255.255.255`, IPv4Reserved},

		// Test
		{false, `191.0.1.255`, IPv4TestNet1},
		{true, `192.0.2.0`, IPv4TestNet1},
		{true, `192.0.2.255`, IPv4TestNet1},
		{false, `192.0.3.0`, IPv4TestNet1},

		{false, `198.50.98.255`, IPv4TestNet2},
		{true, `198.51.100.0`, IPv4TestNet2},
		{true, `198.51.100.255`, IPv4TestNet2},
		{false, `198.51.101.0`, IPv4TestNet2},

		{false, `203.0.112.255`, IPv4TestNet3},
		{true, `203.0.113.0`, IPv4TestNet3},
		{true, `203.0.113.255`, IPv4TestNet3},
		{false, `203.0.114.0`, IPv4TestNet3},

		{false, `233.251.255.255`, IPv4MulticastTest},
		{true, `233.252.0.0`, IPv4MulticastTest},
		{true, `233.252.0.255`, IPv4MulticastTest},
		{false, `233.252.1.0`, IPv4MulticastTest},

		{false, `198.17.255.255`, IPv4NetworkBench},
		{true, `198.18.0.0`, IPv4NetworkBench},
		{true, `198.19.255.255`, IPv4NetworkBench},
		{false, `198.20.0.0`, IPv4NetworkBench},

		// Multicast
		{false, `223.255.255.255`, IPv4Multicast},
		{true, `224.0.0.0`, IPv4Multicast},
		{true, `239.255.255.255`, IPv4Multicast},
		{false, `240.0.0.0`, IPv4Multicast},

		{true, `255.255.255.255`, IPv4LimitedBroadCast},

		// 6to4
		{false, `192.88.98.255`, IPv46to4Relay},
		{true, `192.88.99.0`, IPv46to4Relay},
		{true, `192.88.99.255`, IPv46to4Relay},
		{false, `192.88.100.0`, IPv46to4Relay},
	}

	for _, test := range tests {
		if result := test.target.Contains(net.ParseIP(test.src)); result != test.result {
			t.Errorf(`test result mismatch: ip: %s, should: %v, result:%v`, test.src, test.result, result)
		}
	}
}
