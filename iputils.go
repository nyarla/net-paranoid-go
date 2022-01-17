package paranoid

import (
	"net"
)

func hasPrefix(addr net.IP, prefix []byte) bool {
	if len(prefix) > len(addr) {
		return false
	}
	if len(prefix) > 0 && addr[0] != prefix[0] {
		return false
	}
	if len(prefix) == 1 {
		return true
	}
	if len(prefix) > 1 && addr[1] != prefix[1] {
		return false
	}
	if len(prefix) == 2 {
		return true
	}
	if len(prefix) > 2 && addr[2] != prefix[2] {
		return false
	}
	if len(prefix) == 3 {
		return true
	}
	if len(prefix) > 3 && addr[3] != prefix[3] {
		return false
	}
	if len(prefix) == 4 {
		return true
	}
	if len(prefix) > 4 && addr[4] != prefix[4] {
		return false
	}
	if len(prefix) == 5 {
		return true
	}
	if len(prefix) > 5 && addr[5] != prefix[5] {
		return false
	}
	if len(prefix) == 6 {
		return true
	}
	if len(prefix) > 6 && addr[6] != prefix[6] {
		return false
	}
	if len(prefix) == 7 {
		return true
	}
	if len(prefix) > 7 && addr[7] != prefix[7] {
		return false
	}
	if len(prefix) == 8 {
		return true
	}
	if len(prefix) > 8 && addr[8] != prefix[8] {
		return false
	}
	if len(prefix) == 9 {
		return true
	}
	if len(prefix) > 9 && addr[9] != prefix[9] {
		return false
	}
	if len(prefix) == 10 {
		return true
	}
	if len(prefix) > 10 && addr[10] != prefix[10] {
		return false
	}
	if len(prefix) == 11 {
		return true
	}
	if len(prefix) > 11 && addr[11] != prefix[11] {
		return false
	}
	if len(prefix) == 12 {
		return true
	}
	if len(prefix) > 12 && addr[12] != prefix[12] {
		return false
	}
	if len(prefix) == 13 {
		return true
	}
	if len(prefix) > 13 && addr[13] != prefix[13] {
		return false
	}
	if len(prefix) == 14 {
		return true
	}
	if len(prefix) > 14 && addr[14] != prefix[14] {
		return false
	}
	if len(prefix) == 15 {
		return true
	}
	if addr[15] != prefix[15] {
		return false
	}
	return true
}

var (
	prefixIPv6Loopback       = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	prefixIPv4Inv6           = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xff, 0xff}
	prefixIPv4Inv6Deprecated = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	prefixIPv6Benchmarking   = []byte{0x20, 0x01, 0x0, 0x2, 0x0, 0x0}
	prefixIPv6DiscardOnly    = []byte{1, 0, 0, 0, 0, 0, 0, 0}
	prefixIPv6Teredo         = []byte{0x20, 0x01, 0x00, 0x00}
)

func IsUnspecified(addr net.IP) bool {
	if len(addr) == net.IPv4len {
		return addr[0] == 0 && addr[1] == 0 && addr[2] == 0 && addr[3] == 0
	}

	if len(addr) == net.IPv6len && addr[12] == 0 && addr[13] == 0 && addr[14] == 0 && addr[15] == 0 {
		return (addr[10] == 0xff && addr[11] == 0xff) || (addr[10] == 0 && addr[11] == 0)
	}

	return false
}

func IsPrivate(addr net.IP) bool {
	if len(addr) == net.IPv4len {
		return addr[0] == 10 || (addr[0] == 172 && addr[1]&240 == 16) || (addr[0] == 192 && addr[1] == 168)
	}

	if len(addr) == net.IPv6len {
		if addr[0]&0xfe == 0xfc {
			return true
		}

		if addr[0] == 0xfe && addr[1] >= 0xc0 {
			return true
		}

		return hasPrefix(addr, prefixIPv4Inv6) && (addr[12] == 10 || (addr[12] == 172 && addr[13]&240 == 16) || (addr[12] == 192 && addr[13] == 168))
	}

	return false
}

func IsLoopback(addr net.IP) bool {
	if len(addr) == net.IPv4len {
		return addr[0] == 127
	}

	if len(addr) == net.IPv6len {
		if !hasPrefix(addr, prefixIPv6Loopback) {
			return false
		}

		if addr[10] == addr[11] {
			if addr[10] == 0xff {
				return addr[12] == 127
			}

			return addr[10] == 0 && addr[12] == 0 && addr[13] == 0 && addr[14] == 0 && addr[15] == 1
		}
	}

	return false
}

func IsLinkLocalAddr(addr net.IP) bool {
	if len(addr) == net.IPv4len {
		return addr[0] == 169 && addr[1] == 254
	}

	if len(addr) == net.IPv6len {
		if addr[0] == 0xfe && addr[1]&0xc0 == 0x80 {
			return true
		}

		return hasPrefix(addr, prefixIPv4Inv6) && addr[12] == 169 && addr[13] == 254
	}

	return false
}

func IsSharedAddr(addr net.IP) bool {
	if len(addr) == net.IPv4len {
		return addr[0] == 100 && addr[1]&192 == 64
	}

	if len(addr) == net.IPv6len {
		return hasPrefix(addr, prefixIPv4Inv6) && addr[12] == 100 && addr[13]&192 == 64
	}

	return false
}

func IsProtocolAssignments(addr net.IP) bool {
	if len(addr) == net.IPv4len {
		return addr[0] == 192 && addr[1] == 0 && addr[2] == 0
	}

	if len(addr) == net.IPv6len {
		if addr[0] == 0x20 && addr[1] == 0x01 && addr[2] <= 1 && addr[3] <= 0xff {
			return true
		}

		return hasPrefix(addr, prefixIPv4Inv6) && addr[12] == 192 && addr[13] == 0 && addr[14] == 0
	}

	return false
}

func IsDocuments(addr net.IP) bool {
	if len(addr) == net.IPv4len {
		return (addr[0] == 192 && addr[1] == 0 && addr[2] == 2) || (addr[0] == 198 && addr[1] == 51 && addr[2] == 100) || (addr[0] == 203 && addr[1] == 0 && addr[2] == 113)
	}

	if len(addr) == net.IPv6len {
		if addr[0] == 0x20 && addr[1] == 0x01 && addr[2] == 0x0d && addr[3] == 0xb8 {
			return true
		}

		return hasPrefix(addr, prefixIPv4Inv6) && ((addr[12] == 192 && addr[13] == 0 && addr[14] == 2) || (addr[12] == 198 && addr[13] == 51 && addr[14] == 100) || (addr[12] == 203 && addr[13] == 0 && addr[14] == 113))
	}

	return false
}

func IsIPVersionRelay(addr net.IP) bool {
	if len(addr) == net.IPv4len {
		return addr[0] == 192 && addr[1] == 88 && addr[2] == 99
	}

	if len(addr) == net.IPv6len {
		if addr[0] == 0x20 && addr[1] == 0x1 {
			return true
		}

		return hasPrefix(addr, prefixIPv4Inv6) && addr[12] == 192 && addr[13] == 88 && addr[14] == 99
	}

	return false
}

func IsBenchmarking(addr net.IP) bool {
	if len(addr) == net.IPv4len {
		return addr[0] == 198 && addr[1]&254 == 18
	}

	if len(addr) == net.IPv6len {
		if hasPrefix(addr, prefixIPv6Benchmarking) {
			return true
		}

		return hasPrefix(addr, prefixIPv4Inv6) && addr[12] == 198 && addr[13]&254 == 18
	}

	return false
}

func IsMulticast(addr net.IP) bool {
	if len(addr) == net.IPv4len {
		return addr[0]&240 == 224
	}

	if len(addr) == net.IPv6len {
		if addr[0] == 0xff {
			return true
		}

		return hasPrefix(addr, prefixIPv4Inv6) && addr[12]&240 == 224
	}

	return false
}

func IsReserved(addr net.IP) bool {
	if len(addr) == net.IPv4len {
		return addr[0] >= 240
	}

	if len(addr) == net.IPv6len {
		return hasPrefix(addr, prefixIPv4Inv6) && addr[12] >= 240
	}

	return false
}

func IsDeprecatedIPv4inIPv6(addr net.IP) bool {
	if len(addr) == net.IPv6len {
		return hasPrefix(addr, prefixIPv4Inv6Deprecated)
	}

	return false
}

func IsMigrationIPv6(addr net.IP) bool {
	if len(addr) == net.IPv6len && addr[0] == 0x0 && addr[1] == 0x64 && addr[2] == 0xff && addr[3] == 0x9b {
		if addr[4] == 0 && addr[5] == 1 {
			return true
		}

		return addr[4] == 0 && addr[5] == 0 && addr[6] == 0 && addr[7] == 0 && addr[8] == 0 && addr[9] == 0 && addr[10] == 0 && addr[11] == 0
	}

	return false
}

func IsDiscardOnly(addr net.IP) bool {
	return len(addr) == net.IPv6len && hasPrefix(addr, prefixIPv6DiscardOnly)
}

func IsGlobalUnicast(addr net.IP) bool {
	if len(addr) == net.IPv6len && addr[0]&0x20 == 0x20 {
		return true
	}

	return (len(addr) == net.IPv4len || (len(addr) == net.IPv6len && hasPrefix(addr, prefixIPv4Inv6))) && !(IsUnspecified(addr) || IsLoopback(addr) || IsPrivate(addr) || IsLinkLocalAddr(addr) || IsProtocolAssignments(addr) || IsIPVersionRelay(addr) || IsDocuments(addr) || IsBenchmarking(addr) || IsMulticast(addr) || IsReserved(addr))
}

func IsTeredo(addr net.IP) bool {
	return len(addr) == net.IPv6len && hasPrefix(addr, prefixIPv6Teredo)
}

func IsORCHID(addr net.IP) bool {
	if len(addr) == net.IPv6len && addr[0] == 0x20 && addr[1] == 0x01 && addr[2] == 0x00 {
		return addr[3]&0xb0 == 0x10 || addr[3]&0xb0 == 0x20
	}

	return false
}

func IsBlockRecommended(addr net.IP) bool {
	return IsTeredo(addr) || IsORCHID(addr) || !IsGlobalUnicast(addr)
}
