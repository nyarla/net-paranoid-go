package paranoid

import (
	"context"
	"net"
	"sync"
	"testing"
	"time"
)

func TestParanoidDialer(t *testing.T) {
	parent := new(net.Dialer)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)

	ip := NewBlockIPRules(IsBlockRecommended)
	host := NewAllowHostRules(StringHostMatcher(`the.kalaclista.com`, IsSameHost))

	defer cancel()

	dialer := NewDialer(parent, ip, host)

	if _, err := dialer.DialContext(ctx, `tcp`, `the.kalaclista.com:443`); err != nil {
		t.Errorf(`failed to test for ParanodDialer: %v`, err)
	}
}

func BenchmarkParanoidDialer(b *testing.B) {
	parent := new(net.Dialer)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)

	ip := NewBlockIPRules(IsBlockRecommended)
	host := NewAllowHostRules(StringHostMatcher(`the.kalaclista.com`, IsSameHost))

	defer cancel()
	var once sync.Once

	dialer := NewDialer(parent, ip, host)

	b.ReportAllocs()
	once.Do(func() {
		dialer.DialContext(ctx, `tcp`, `the.kalaclista.com:443`)
	})
}
