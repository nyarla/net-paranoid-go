package paranoid

import (
	"net"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

func TestIPUtils(t *testing.T) {
	tests := []struct {
		fn   func(addr net.IP) bool
		src  string
		want bool
	}{
		// Unspecified
		{IsUnspecified, `0.0.0.0`, true},
		{IsUnspecified, `::`, true},

		{IsUnspecified, `127.0.0.1`, false},
		{IsUnspecified, `::1`, false},

		// Loopback
		{IsLoopback, `127.0.0.0`, true},
		{IsLoopback, `127.255.255.255`, true},
		{IsLoopback, `::1`, true},

		{IsLoopback, `126.255.255.255`, false},
		{IsLoopback, `128.0.0.0`, false},
		{IsLoopback, `::2`, false},

		// Private Use
		{IsPrivate, `10.0.0.0`, true},
		{IsPrivate, `10.255.255.255`, true},

		{IsPrivate, `9.255.255.255`, false},
		{IsPrivate, `11.0.0.0`, false},

		{IsPrivate, `172.16.0.0`, true},
		{IsPrivate, `172.31.255.255`, true},
		{IsPrivate, `fc00::`, true},
		{IsPrivate, `fdff:ffff:ffff:ffff:ffff:ffff:ffff:ffff`, true},
		{IsPrivate, `fec0::`, true},
		{IsPrivate, `feff:ffff:ffff:ffff:ffff:ffff:ffff:ffff`, true},

		{IsPrivate, `172.15.255.255`, false},
		{IsPrivate, `172.32.0.0`, false},
		{IsPrivate, `fbff:ffff:ffff:ffff:ffff:ffff:ffff:ffff`, false},
		{IsPrivate, `fe00::`, false},
		{IsPrivate, `febf:ffff:ffff:ffff:ffff:ffff:ffff:ffff`, false},
		{IsPrivate, `ff00::`, false},

		// Shared Address (IPv4)
		{IsSharedAddr, `100.64.0.0`, true},
		{IsSharedAddr, `100.127.255.255`, true},

		{IsSharedAddr, `100.63.255.255`, false},
		{IsSharedAddr, `100.128.0.0`, false},

		// Link Local
		{IsLinkLocalAddr, `169.254.0.0`, true},
		{IsLinkLocalAddr, `169.254.255.255`, true},
		{IsLinkLocalAddr, `fe80::`, true},
		{IsLinkLocalAddr, `febf:ffff:ffff:ffff:ffff:ffff:ffff:ffff`, true},

		{IsLinkLocalAddr, `169.253.255.255`, false},
		{IsLinkLocalAddr, `169.255.0.0`, false},
		{IsLinkLocalAddr, `fe7f:ffff:ffff:ffff:ffff:ffff:ffff:ffff`, false},
		{IsLinkLocalAddr, `fec0::`, false},

		// IETF Protocol Assignments
		{IsProtocolAssignments, `192.0.0.0`, true},
		{IsProtocolAssignments, `192.0.0.255`, true},
		{IsProtocolAssignments, `2001::`, true},
		{IsProtocolAssignments, `2001:01ff:ffff:ffff:ffff:ffff:ffff:ffff`, true},

		{IsProtocolAssignments, `191.255.255.255`, false},
		{IsProtocolAssignments, `192.0.1.0`, false},
		{IsProtocolAssignments, `2000:ffff:ffff:ffff:ffff:ffff:ffff:ffff`, false},
		{IsProtocolAssignments, `2001:0200::`, false},

		// Documents
		{IsDocuments, `192.0.2.0`, true},
		{IsDocuments, `192.0.2.255`, true},
		{IsDocuments, `198.51.100.0`, true},
		{IsDocuments, `198.51.100.255`, true},
		{IsDocuments, `203.0.113.0`, true},
		{IsDocuments, `203.0.113.255`, true},
		{IsDocuments, `2001:db8::`, true},
		{IsDocuments, `2001:db8:ffff:ffff:ffff:ffff:ffff:ffff`, true},

		{IsDocuments, `192.0.1.255`, false},
		{IsDocuments, `192.0.3.0`, false},
		{IsDocuments, `198.51.99.255`, false},
		{IsDocuments, `198.52.101.0`, false},
		{IsDocuments, `203.0.112.225`, false},
		{IsDocuments, `203.0.114.0`, false},
		{IsDocuments, `2001:db7:ffff:ffff:ffff:ffff:ffff:ffff`, false},
		{IsDocuments, `2001:db9::`, false},

		// 4to6 or 6to4 relay
		{IsIPVersionRelay, `192.88.99.0`, true},
		{IsIPVersionRelay, `192.88.99.255`, true},
		{IsIPVersionRelay, `2001::`, true},
		{IsIPVersionRelay, `2001:ffff:ffff:ffff:ffff:ffff:ffff:ffff`, true},

		{IsIPVersionRelay, `192.88.98.255`, false},
		{IsIPVersionRelay, `192.88.100.0`, false},
		{IsIPVersionRelay, `2000:ffff:ffff:ffff:ffff:ffff:ffff:ffff`, false},
		{IsIPVersionRelay, `2002::`, false},

		// For benchmarking
		{IsBenchmarking, `198.18.0.0`, true},
		{IsBenchmarking, `198.19.255.255`, true},
		{IsBenchmarking, `2001:2::`, true},
		{IsBenchmarking, `2001:2:0:ffff:ffff:ffff:ffff:ffff`, true},

		{IsBenchmarking, `198.17.255.255`, false},
		{IsBenchmarking, `192.20.0.0`, false},
		{IsBenchmarking, `2001:1:ffff:ffff:ffff:ffff:ffff:ffff`, false},
		{IsBenchmarking, `2001:2:1::`, false},

		// Multicast
		{IsMulticast, `224.0.0.0`, true},
		{IsMulticast, `239.255.255.255`, true},
		{IsMulticast, `ff00::`, true},
		{IsMulticast, `ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff`, true},

		{IsMulticast, `223.255.255.255`, false},
		{IsMulticast, `240.255.255.255`, false},
		{IsMulticast, `feff:ffff:ffff:ffff:ffff:ffff:ffff:ffff`, false},

		// Reserved
		{IsReserved, `240.0.0.0`, true},
		{IsReserved, `255.255.255.255`, true},

		{IsReserved, `239.255.255.255`, false},

		// Deprecated IPv4 in IPv6
		{IsDeprecatedIPv4inIPv6, `::0`, true},
		{IsDeprecatedIPv4inIPv6, `::ffff:ffff`, true},

		{IsDeprecatedIPv4inIPv6, `::1:0:0`, false},

		// IPv6 Migration
		{IsMigrationIPv6, `64:ff9b::`, true},
		{IsMigrationIPv6, `64:ff9b::ffff:ffff`, true},
		{IsMigrationIPv6, `64:ff9b:1::`, true},
		{IsMigrationIPv6, `64:ff9b:1:ffff:ffff:ffff:ffff:ffff`, true},

		{IsMigrationIPv6, `64:ff9a:ffff:ffff:ffff:ffff:ffff:ffff`, false},
		{IsMigrationIPv6, `64:ff9b::1:0:0`, false},
		{IsMigrationIPv6, `64:ff9a:0:ffff:ffff:ffff:ffff:ffff`, false},
		{IsMigrationIPv6, `64:ff9b:2::`, false},

		// IPv6 Discard only
		{IsDiscardOnly, `100::`, true},
		{IsDiscardOnly, `100::ffff:ffff:ffff:ffff`, true},

		{IsDiscardOnly, `99:ffff:ffff:ffff:ffff:ffff:ffff:ffff`, false},
		{IsDiscardOnly, `100::1:0:0:0:0`, false},

		// IPv6 GlobalUnicat
		{IsGlobalUnicast, `1.0.0.0`, true},
		{IsGlobalUnicast, `100.63.255.255`, true},
		{IsGlobalUnicast, `100.65.0.0`, true},
		{IsGlobalUnicast, `126.255.255.255`, true},
		{IsGlobalUnicast, `172.15.255.255`, true},
		{IsGlobalUnicast, `172.32.0.0`, true},
		{IsGlobalUnicast, `191.255.255.255`, true},
		{IsGlobalUnicast, `192.0.1.255`, true},
		{IsGlobalUnicast, `192.0.3.0`, true},
		{IsGlobalUnicast, `192.88.98.255`, true},
		{IsGlobalUnicast, `192.88.100.0`, true},
		{IsGlobalUnicast, `192.167.255.255`, true},
		{IsGlobalUnicast, `192.169.0.0`, true},
		{IsGlobalUnicast, `198.17.255.255`, true},
		{IsGlobalUnicast, `198.20.0.0`, true},
		{IsGlobalUnicast, `198.51.99.255`, true},
		{IsGlobalUnicast, `198.51.101.0`, true},
		{IsGlobalUnicast, `203.0.112.255`, true},
		{IsGlobalUnicast, `203.0.114.0`, true},

		{IsGlobalUnicast, `2000::`, true},
		{IsGlobalUnicast, `2001:db7:ffff:ffff:ffff:ffff:ffff:ffff`, true},
		{IsGlobalUnicast, `2001:db9::`, true},
		{IsGlobalUnicast, `3fff:ffff:ffff:ffff:ffff:ffff:ffff:ffff`, true},

		{IsGlobalUnicast, `1fff:ffff:ffff:ffff:ffff:ffff:ffff:ffff`, false},
		{IsGlobalUnicast, `4000::`, false},

		// Teredo
		{IsTeredo, `2001::`, true},
		{IsTeredo, `2001:0:ffff:ffff:ffff:ffff:ffff:ffff`, true},

		{IsTeredo, `2000:ffff:ffff:ffff:ffff:ffff:ffff:ffff`, false},
		{IsTeredo, `2000:1::`, false},

		// ORCHID
		{IsORCHID, `2001:10::`, true},
		{IsORCHID, `2001:2f:ffff:ffff:ffff:ffff:ffff:ffff`, true},

		{IsORCHID, `2001:f:ffff:ffff:ffff:ffff:ffff:ffff`, false},
		{IsORCHID, `2001:30::`, false},
	}

	for _, test := range tests {
		// IPv4 test
		fname := strings.Split(
			(runtime.FuncForPC(reflect.ValueOf(test.fn).Pointer()).Name()), `.`,
		)

		addr := net.ParseIP(test.src)

		if strings.Contains(test.src, `.`) {
			if result := test.fn(addr.To4()); result != test.want {
				t.Errorf("Test .%s failed (IPv4):\naddr: %s,\nwant: %t,\ngot: %t",
					fname[len(fname)-1], test.src, test.want, result,
				)
			}
		}

		if result := test.fn(addr.To16()); result != test.want {
			t.Errorf("Test .%s failed (IPv6):\naddr: %s,\nwant: %t,\ngot: %t",
				fname[len(fname)-1], test.src, test.want, result,
			)
		}
	}
}
