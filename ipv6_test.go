package paranoid

import (
	"net"
	"testing"
)

func TestIPv6(t *testing.T) {
	type task struct {
		result bool
		src    string
		target *net.IPNet
	}

	tests := []task{
		{false, `::`, IPv6Loopback},
		{true, `::1`, IPv6Loopback},
		{false, `::2`, IPv6Loopback},

		{true, `::`, IPv6DeprecatedIPv4Compatible},
		{true, `::1`, IPv6DeprecatedIPv4Compatible},
		{true, `::ffff:ffff`, IPv6DeprecatedIPv4Compatible},
		{false, `::1:0:0`, IPv6DeprecatedIPv4Compatible},

		{false, `::fffe:0:0`, IPv6MappingIPv4},
		{true, `::ffff:0:1`, IPv6MappingIPv4},
		{true, `::ffff:ffff:ffff`, IPv6MappingIPv4},
		{false, `:1:0:0:0`, IPv6MappingIPv4},

		{false, `64:ff9a:ffff:ffff:ffff:ffff:ffff:ffff`, IPv6MigrationRFC6052},
		{true, `64:ff9b::`, IPv6MigrationRFC6052},
		{true, `64:ff9b::ffff:ffff`, IPv6MigrationRFC6052},
		{false, `64:ff9c::1:0:0`, IPv6MigrationRFC6052},

		{false, `64:ff9b:0:ffff:ffff:ffff:ffff:ffff`, IPv6MigrationRFC8215},
		{true, `64:ff9b:1::`, IPv6MigrationRFC8215},
		{true, `64:ff9b:1:ffff:ffff:ffff:ffff:ffff`, IPv6MigrationRFC8215},
		{false, `64:ff9c:2::`, IPv6MigrationRFC8215},

		{false, `99:ffff:ffff:ffff:ffff:ffff:ffff:ffff`, IPv6DiscardOnly},
		{true, `100::`, IPv6DiscardOnly},
		{true, `100::ffff:ffff:ffff:ffff`, IPv6DiscardOnly},
		{false, `100::1:0:0:0:0`, IPv6DiscardOnly},

		{false, `199f:ffff:ffff:ffff:ffff:ffff:ffff:ffff`, IPv6GlobalUnicast},
		{true, `2000::`, IPv6GlobalUnicast},
		{true, `3fff:ffff:ffff:ffff:ffff:ffff:ffff:ffff`, IPv6GlobalUnicast},
		{false, `4000::`, IPv6GlobalUnicast},

		{false, `2000:ffff:ffff:ffff:ffff:ffff:ffff:ffff`, IPv6SubTLAIDAssign},
		{true, `2001::`, IPv6SubTLAIDAssign},
		{true, `2001:001f:ffff:ffff:ffff:ffff:ffff:ffff`, IPv6SubTLAIDAssign},
		{false, `2001:0200::`, IPv6SubTLAIDAssign},

		{false, `2000:ffff:ffff:ffff:ffff:ffff:ffff:ffff`, IPv6Teredo},
		{true, `2001::`, IPv6Teredo},
		{true, `2001:0:ffff:ffff:ffff:ffff:ffff:ffff`, IPv6Teredo},
		{false, `2001:1::`, IPv6Teredo},

		{false, `2001:1:ffff:ffff:ffff:ffff:ffff:ffff`, IPv6Benchmarking},
		{true, `2001:2::`, IPv6Benchmarking},
		{true, `2001:2:0:ffff:ffff:ffff:ffff:ffff`, IPv6Benchmarking},
		{false, `2001:2:1::`, IPv6Benchmarking},

		{false, `2001:2:ffff:ffff:ffff:ffff:ffff:ffff`, IPv6AutomaticMulticastTunneling},
		{true, `2001:3::`, IPv6AutomaticMulticastTunneling},
		{true, `2001:3:ffff:ffff:ffff:ffff:ffff:ffff`, IPv6AutomaticMulticastTunneling},
		{false, `2001:4::`, IPv6AutomaticMulticastTunneling},

		{false, `2001:4:111:ffff:ffff:ffff:ffff:ffff`, IPv6AS112RedirectionUsingDNAME},
		{true, `2001:4:112::`, IPv6AS112RedirectionUsingDNAME},
		{true, `2001:4:112:ffff:ffff:ffff:ffff:ffff`, IPv6AS112RedirectionUsingDNAME},
		{false, `2001:4:113::`, IPv6AS112RedirectionUsingDNAME},

		{false, `2001:4:ffff:ffff:ffff:ffff:ffff:ffff`, IPv6LISPEID},
		{true, `2001:5::`, IPv6LISPEID},
		{true, `2001:5:ffff:ffff:ffff:ffff:ffff:ffff`, IPv6LISPEID},
		{false, `2001:6::`, IPv6LISPEID},

		{false, `2001:f:ffff:ffff:ffff:ffff:ffff:ffff`, IPv6ORCHIDDeprecated},
		{true, `2001:10::`, IPv6ORCHIDDeprecated},
		{true, `2001:1f:ffff:ffff:ffff:ffff:ffff:ffff`, IPv6ORCHIDDeprecated},
		{false, `2001:20::`, IPv6ORCHIDDeprecated},

		{false, `2001:1f:ffff:ffff:ffff:ffff:ffff:ffff`, IPv6ORCHIDv2},
		{true, `2001:20::`, IPv6ORCHIDv2},
		{true, `2001:2f:ffff:ffff:ffff:ffff:ffff:ffff`, IPv6ORCHIDv2},
		{false, `2001:30::`, IPv6ORCHIDv2},

		{false, `2001:db7:ffff:ffff:ffff:ffff:ffff:ffff`, IPv6Documents},
		{true, `2001:db8::`, IPv6Documents},
		{true, `2001:db8:ffff:ffff:ffff:ffff:ffff:ffff`, IPv6Documents},
		{false, `2001:db9::`, IPv6Documents},

		{false, `2001:ffff:ffff:ffff:ffff:ffff:ffff:ffff`, IPv6Deprecated6to4AnyCast},
		{true, `2002::`, IPv6Deprecated6to4AnyCast},
		{true, `2002:ffff:ffff:ffff:ffff:ffff:ffff:ffff`, IPv6Deprecated6to4AnyCast},
		{false, `2003::`, IPv6Deprecated6to4AnyCast},

		{false, `2620:4f:7fff:ffff:ffff:ffff:ffff:ffff`, IPv6AS112NameserverOperations},
		{true, `2620:4f:8000::`, IPv6AS112NameserverOperations},
		{true, `2620:4f:8000:ffff:ffff:ffff:ffff:ffff`, IPv6AS112NameserverOperations},
		{false, `2620:4f:8001::`, IPv6AS112NameserverOperations},

		{false, `3ffd:ffff:ffff:ffff:ffff:ffff:ffff:ffff`, IPv6Bone},
		{true, `3ffe::`, IPv6Bone},
		{true, `3ffe:ffff:ffff:ffff:ffff:ffff:ffff:ffff`, IPv6Bone},
		{false, `3fff::`, IPv6Bone},

		{false, `fbff:ffff:ffff:ffff:ffff:ffff:ffff:ffff`, IPv6UniqueLocalUnicast},
		{true, `fc00::`, IPv6UniqueLocalUnicast},
		{true, `fdff:ffff:ffff:ffff:ffff:ffff:ffff:ffff`, IPv6UniqueLocalUnicast},
		{false, `fe00::`, IPv6UniqueLocalUnicast},

		{false, `fbff:ffff:ffff:ffff:ffff:ffff:ffff:ffff`, IPv6AuthorityManaged},
		{true, `fc00::`, IPv6AuthorityManaged},
		{true, `fcff:ffff:ffff:ffff:ffff:ffff:ffff:ffff`, IPv6AuthorityManaged},
		{false, `fd00::`, IPv6AuthorityManaged},

		{false, `fcff:ffff:ffff:ffff:ffff:ffff:ffff:ffff`, IPv6LocalAssigned},
		{true, `fd00::`, IPv6LocalAssigned},
		{true, `fdff:ffff:ffff:ffff:ffff:ffff:ffff:ffff`, IPv6LocalAssigned},
		{false, `fe00::`, IPv6LocalAssigned},

		{false, `fe7f:ffff:ffff:ffff:ffff:ffff:ffff:ffff`, IPv6LinkLocalUnicast},
		{true, `fe80::`, IPv6LinkLocalUnicast},
		{true, `febf:ffff:ffff:ffff:ffff:ffff:ffff:ffff`, IPv6LinkLocalUnicast},
		{false, `fec0::`, IPv6LinkLocalUnicast},

		{false, `febf:ffff:ffff:ffff:ffff:ffff:ffff:ffff`, IPv6SiteLocalUnicast},
		{true, `fec0::`, IPv6SiteLocalUnicast},
		{true, `feff:ffff:ffff:ffff:ffff:ffff:ffff:ffff`, IPv6SiteLocalUnicast},
		{false, `ff00::`, IPv6SiteLocalUnicast},

		{false, `feff:ffff:ffff:ffff:ffff:ffff:ffff:ffff`, IPv6Multicast},
		{true, `ff00::`, IPv6Multicast},
		{true, `ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff`, IPv6Multicast},

		{false, `ff00:ffff:ffff:ffff:ffff:ffff:ffff:ffff`, IPv6NodeLocalMulticast},
		{true, `ff01::`, IPv6NodeLocalMulticast},
		{true, `ff01:ffff:ffff:ffff:ffff:ffff:ffff:ffff`, IPv6NodeLocalMulticast},
		{false, `ff02::`, IPv6NodeLocalMulticast},

		{false, `ff01:ffff:ffff:ffff:ffff:ffff:ffff:ffff`, IPv6LinkLocalMulticast},
		{true, `ff02::`, IPv6LinkLocalMulticast},
		{true, `ff02:ffff:ffff:ffff:ffff:ffff:ffff:ffff`, IPv6LinkLocalMulticast},
		{false, `ff03::`, IPv6LinkLocalMulticast},

		{false, `ff04:ffff:ffff:ffff:ffff:ffff:ffff:ffff`, IPv6SiteLocalMulticast},
		{true, `ff05::`, IPv6SiteLocalMulticast},
		{true, `ff05:ffff:ffff:ffff:ffff:ffff:ffff:ffff`, IPv6SiteLocalMulticast},
		{false, `ff06::`, IPv6SiteLocalMulticast},

		{false, `ff0d:ffff:ffff:ffff:ffff:ffff:ffff:ffff`, IPv6GlobalMulticast},
		{true, `ff0e::`, IPv6GlobalMulticast},
		{true, `ff0e:ffff:ffff:ffff:ffff:ffff:ffff:ffff`, IPv6GlobalMulticast},
		{false, `ff0f::`, IPv6GlobalMulticast},
	}

	for _, test := range tests {
		if result := test.target.Contains(net.ParseIP(test.src)); result != test.result {
			t.Errorf(`test result mismatch: ip: %s, should: %v, result:%v`, test.src, test.result, result)
		}
	}
}
