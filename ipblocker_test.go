package paranoid

import (
	"crypto/rand"
	"net"
	"testing"
)

func TestIPBlocker(t *testing.T) {
	var (
		list    []*net.IPNet
		err     error
		blocker *IPBlocker
		ip      net.IP
	)

	list = make([]*net.IPNet, 0)
	if _, err = NewIPBlocker(list, true); err != ErrIPNetMissing {
		t.Fatal("In this case, this function must throw ErrIPNetMissing error.")
	}

	list = make([]*net.IPNet, 1)
	list[0] = IPv4PrivateClassA

	if blocker, err = NewIPBlocker(list, true); err != nil {
		t.Fatal("In this case, this function should not return error.")
	}

	ip = net.IPv4(10, 0, 0, 1)
	if !blocker.IsForbiddenIP(ip) {
		t.Fatal("In this case, this IP (10.0.0.1) should be block")
	}

	if blocker, err = NewIPBlocker(list, false); err != nil {
		t.Fatal("In this case, this function should not return error.")
	}

	if blocker.IsForbiddenIP(ip) {
		t.Fatal("In this case, this IP (10.0.0.1) should not be block")
	}
}

func BenchmarkIPBlocker(b *testing.B) {
	var (
		src []byte = make([]byte, 4)
		ips        = make([]net.IP, 64)
	)

	blocker, _ := NewIPBlocker(IPv4RangesBlockRecommended, true)

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
