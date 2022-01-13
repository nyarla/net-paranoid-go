package paranoid

import (
	"net"
)

var (
	// List of IPv4 blocks
	IPv4PrivateClassA    = &net.IPNet{IP: net.IPv4(10, 0, 0, 0), Mask: net.CIDRMask(8, 32)}
	IPv4SharedAddr       = &net.IPNet{IP: net.IPv4(100, 64, 0, 0), Mask: net.CIDRMask(10, 32)}
	IPv4Loopback         = &net.IPNet{IP: net.IPv4(127, 0, 0, 0), Mask: net.CIDRMask(8, 32)}
	IPv4LinkLocal        = &net.IPNet{IP: net.IPv4(169, 254, 0, 0), Mask: net.CIDRMask(16, 32)}
	IPv4PrivateClassB    = &net.IPNet{IP: net.IPv4(172, 16, 0, 0), Mask: net.CIDRMask(12, 32)}
	IPv4SpecialPurpose   = &net.IPNet{IP: net.IPv4(192, 0, 0, 0), Mask: net.CIDRMask(24, 32)}
	IPv4TestNet1         = &net.IPNet{IP: net.IPv4(192, 0, 2, 0), Mask: net.CIDRMask(24, 32)}
	IPv46to4Relay        = &net.IPNet{IP: net.IPv4(192, 88, 99, 0), Mask: net.CIDRMask(24, 32)}
	IPv4PrivateClassC    = &net.IPNet{IP: net.IPv4(192, 168, 0, 0), Mask: net.CIDRMask(16, 32)}
	IPv4NetworkBench     = &net.IPNet{IP: net.IPv4(198, 18, 0, 0), Mask: net.CIDRMask(15, 32)}
	IPv4TestNet2         = &net.IPNet{IP: net.IPv4(198, 51, 100, 0), Mask: net.CIDRMask(24, 32)}
	IPv4TestNet3         = &net.IPNet{IP: net.IPv4(203, 0, 113, 0), Mask: net.CIDRMask(24, 32)}
	IPv4Multicast        = &net.IPNet{IP: net.IPv4(224, 0, 0, 0), Mask: net.CIDRMask(4, 32)}
	IPv4MulticastTest    = &net.IPNet{IP: net.IPv4(233, 252, 0, 0), Mask: net.CIDRMask(24, 32)}
	IPv4Reserved         = &net.IPNet{IP: net.IPv4(240, 0, 0, 0), Mask: net.CIDRMask(4, 32)}
	IPv4LimitedBroadCast = &net.IPNet{IP: net.IPv4(255, 255, 255, 255), Mask: net.CIDRMask(32, 32)}
)

var (
	// Pre-defined IPv4 block ranges
	IPv4RangesPrivate          = []*net.IPNet{IPv4PrivateClassA, IPv4PrivateClassB, IPv4PrivateClassC}
	IPv4RangesInternal         = []*net.IPNet{IPv4SharedAddr, IPv4Loopback, IPv4LinkLocal}
	IPv4RangesReserved         = []*net.IPNet{IPv4SpecialPurpose, IPv4Reserved}
	IPv4RangesTest             = []*net.IPNet{IPv4TestNet1, IPv4TestNet2, IPv4TestNet3, IPv4MulticastTest, IPv4NetworkBench}
	IPv4RangesMultiCast        = []*net.IPNet{IPv4Multicast, IPv4LimitedBroadCast}
	IPv4RangesBlockRecommended = []*net.IPNet{
		IPv4PrivateClassA,
		IPv4PrivateClassB,
		IPv4PrivateClassC,

		IPv4SharedAddr,
		IPv4Loopback,
		IPv4LinkLocal,

		IPv4SpecialPurpose,
		IPv4Reserved,

		IPv4TestNet1,
		IPv4TestNet2,
		IPv4TestNet3,
		IPv4MulticastTest,
		IPv4NetworkBench,

		IPv4Multicast,
		IPv4LimitedBroadCast,

		IPv46to4Relay,
	}
)
