package paranoid

import (
	"crypto/rand"
	"net"
	"testing"
)

func TestIRanges(t *testing.T) {
	blocker := NewBlockIPRunes(IsBlockRecommended)
	allower := NewAllowIPRunes(IsPrivate)
	ip := net.IPv4(10, 0, 0, 1)

	if !blocker.IsForbiddenIP(ip) {
		t.Errorf("Test IPRunes.IsForbiddenIP failed:\nwant: true\ngot: false")
	}

	if allower.IsForbiddenIP(ip) {
		t.Errorf("Test IPRunes.IsForbiddenIP failed:\nwant: false\ngot: true")
	}
}

func BenchmarkIPRunes(b *testing.B) {
	var (
		src []byte = make([]byte, 4)
		ips        = make([]net.IP, 64)
	)

	blocker := NewBlockIPRunes(IsBlockRecommended)

	for idx := 0; idx < 64; idx++ {
		rand.Read(src)
		ips[idx] = net.IPv4(src[0], src[1], src[2], src[3])
	}

	b.ReportAllocs()
	b.ResetTimer()

	for count := 0; count < b.N; count++ {
		for idx := 0; idx < 64; idx++ {
			blocker.IsForbiddenIP(ips[idx])
		}
	}
}

func BenchmarkHostRule(b *testing.B) {
	var (
		hostname = `localhost`
		same     = NewHostRule(IsSameHost, `localhost`)
		prefix   = NewHostRule(HostHasPrefix, `cdn`)
		suffix   = NewHostRule(HostHasSuffix, `example.com`)
	)

	b.ReportAllocs()
	b.ResetTimer()

	for count := 0; count < b.N; count++ {
		same.IsForbiddenHost(hostname)
		prefix.IsForbiddenHost(hostname)
		suffix.IsForbiddenHost(hostname)
	}

}
