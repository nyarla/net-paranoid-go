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

var (
	// Lsit of IPv6 blocks
	IPv6Loopback                    = &net.IPNet{IP: net.ParseIP(`::1`), Mask: net.CIDRMask(128, 128)}
	IPv6DeprecatedIPv4Compatible    = &net.IPNet{IP: net.ParseIP(`::`), Mask: net.CIDRMask(96, 128)}
	IPv6MappingIPv4                 = &net.IPNet{IP: net.ParseIP(`::ffff:0:0`), Mask: net.CIDRMask(96, 128)}
	IPv6MigrationRFC6052            = &net.IPNet{IP: net.ParseIP(`64:ff9b::`), Mask: net.CIDRMask(96, 128)}
	IPv6MigrationRFC8215            = &net.IPNet{IP: net.ParseIP(`64:ff9b:1::`), Mask: net.CIDRMask(48, 128)}
	IPv6DiscardOnly                 = &net.IPNet{IP: net.ParseIP(`100::`), Mask: net.CIDRMask(64, 128)}
	IPv6GlobalUnicast               = &net.IPNet{IP: net.ParseIP(`2000::`), Mask: net.CIDRMask(3, 128)}
	IPv6SubTLAIDAssign              = &net.IPNet{IP: net.ParseIP(`2001::`), Mask: net.CIDRMask(23, 128)}
	IPv6Teredo                      = &net.IPNet{IP: net.ParseIP(`2001::`), Mask: net.CIDRMask(32, 128)}
	IPv6Benchmarking                = &net.IPNet{IP: net.ParseIP(`2001:2::`), Mask: net.CIDRMask(48, 128)}
	IPv6AutomaticMulticastTunneling = &net.IPNet{IP: net.ParseIP(`2001:3::`), Mask: net.CIDRMask(32, 128)}
	IPv6AS112RedirectionUsingDNAME  = &net.IPNet{IP: net.ParseIP(`2001:4:112::`), Mask: net.CIDRMask(48, 128)}
	IPv6LISPEID                     = &net.IPNet{IP: net.ParseIP(`2001:5::`), Mask: net.CIDRMask(32, 128)}
	IPv6ORCHIDDeprecated            = &net.IPNet{IP: net.ParseIP(`2001:10::`), Mask: net.CIDRMask(28, 128)}
	IPv6ORCHIDv2                    = &net.IPNet{IP: net.ParseIP(`2001:20::`), Mask: net.CIDRMask(28, 128)}
	IPv6Documents                   = &net.IPNet{IP: net.ParseIP(`2001:db8::`), Mask: net.CIDRMask(32, 128)}
	IPv6Deprecated6to4AnyCast       = &net.IPNet{IP: net.ParseIP(`2002::`), Mask: net.CIDRMask(16, 128)}
	IPv6AS112NameserverOperations   = &net.IPNet{IP: net.ParseIP(`2620:4f:8000::`), Mask: net.CIDRMask(48, 128)}
	IPv6Bone                        = &net.IPNet{IP: net.ParseIP(`3ffe::`), Mask: net.CIDRMask(16, 128)}
	IPv6UniqueLocalUnicast          = &net.IPNet{IP: net.ParseIP(`fc00::`), Mask: net.CIDRMask(7, 128)}
	IPv6AuthorityManaged            = &net.IPNet{IP: net.ParseIP(`fc00::`), Mask: net.CIDRMask(8, 128)}
	IPv6LocalAssigned               = &net.IPNet{IP: net.ParseIP(`fd00::`), Mask: net.CIDRMask(8, 128)}
	IPv6LinkLocalUnicast            = &net.IPNet{IP: net.ParseIP(`fe80::`), Mask: net.CIDRMask(10, 128)}
	IPv6SiteLocalUnicast            = &net.IPNet{IP: net.ParseIP(`fec0::`), Mask: net.CIDRMask(10, 128)}
	IPv6Multicast                   = &net.IPNet{IP: net.ParseIP(`ff00::`), Mask: net.CIDRMask(8, 128)}
	IPv6NodeLocalMulticast          = &net.IPNet{IP: net.ParseIP(`ff01::`), Mask: net.CIDRMask(16, 128)}
	IPv6LinkLocalMulticast          = &net.IPNet{IP: net.ParseIP(`ff02::`), Mask: net.CIDRMask(16, 128)}
	IPv6SiteLocalMulticast          = &net.IPNet{IP: net.ParseIP(`ff05::`), Mask: net.CIDRMask(16, 128)}
	IPv6GlobalMulticast             = &net.IPNet{IP: net.ParseIP(`ff0e::`), Mask: net.CIDRMask(16, 128)}
)

var (
	// Pre-defined IPv6 block ranges
	IPv6BlockRecommended = []*net.IPNet{
		{IP: net.ParseIP(`::`), Mask: net.CIDRMask(3, 128)},
		{IP: net.ParseIP(`4000::`), Mask: net.CIDRMask(2, 128)},
		{IP: net.ParseIP(`8000::`), Mask: net.CIDRMask(1, 128)},

		IPv6Teredo,
		IPv6ORCHIDDeprecated,
		IPv6ORCHIDv2,
		IPv6Documents,
		IPv6Deprecated6to4AnyCast,
	}
)

var (
	// Pre-defined block recommendation about IPv4 and IPv6
	IPBlockRecommended = []*net.IPNet{
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

		{IP: net.ParseIP(`::`), Mask: net.CIDRMask(3, 128)},
		{IP: net.ParseIP(`4000::`), Mask: net.CIDRMask(2, 128)},
		{IP: net.ParseIP(`8000::`), Mask: net.CIDRMask(1, 128)},

		IPv6Teredo,
		IPv6ORCHIDDeprecated,
		IPv6ORCHIDv2,
		IPv6Documents,
		IPv6Deprecated6to4AnyCast,
	}
)
