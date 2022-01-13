package paranoid

import "net"

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
