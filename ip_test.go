package paranoid

import (
	"crypto/rand"
	"net"
	"testing"
)

func TestIRules(t *testing.T) {
	blocker := NewBlockIPRules(IsBlockRecommended)
	allower := NewAllowIPRules(IsPrivate)
	ip := net.IPv4(10, 0, 0, 1)

	if !blocker.IsForbiddenIP(ip) {
		t.Errorf("Test IPRules.IsForbiddenIP failed:\nwant: true\ngot: false")
	}

	if allower.IsForbiddenIP(ip) {
		t.Errorf("Test IPRules.IsForbiddenIP failed:\nwant: false\ngot: true")
	}
}

func BenchmarkIPRules(b *testing.B) {
	var (
		src []byte = make([]byte, 4)
		ips        = make([]net.IP, 64)
	)

	blocker := NewBlockIPRules(IsBlockRecommended)

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
