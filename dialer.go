package paranoid

import (
	"context"
	"fmt"
	"net"
	"sync"
)

type ParanoidDialer interface {
	DialContext(ctx context.Context, network, addr string) (net.Conn, error)
}

type paranoidDialer struct {
	ipRules   IPRules
	hostRules HostRules
	parent    *net.Dialer
	mutex     sync.Mutex
	idx       int
	ip        net.IP
	host      string
	port      string
	addrs     []net.IPAddr
	err       error
	resolver  *net.Resolver
}

func (this *paranoidDialer) DialContext(ctx context.Context, network, addr string) (net.Conn, error) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	switch network {
	case "tcp", "tcp4", "tcp6", "udp", "udp4", "udp6":
	default:
		this.err = fmt.Errorf(`does not support any network except tcp or udp`)
		return nil, this.err
	}

	this.host, this.port, this.err, this.ip, this.addrs, this.idx = ``, ``, nil, nil, nil, 0

	this.host, this.port, this.err = net.SplitHostPort(addr)
	if this.err != nil {
		return nil, this.err
	}

	if this.ip = net.ParseIP(this.host); this.ip != nil {
		if this.ipRules.IsForbiddenIP(this.ip) {
			this.err = fmt.Errorf(`this ip exists on block list: %s`, this.host)
			return nil, this.err
		}
	}

	if this.hostRules.IsForbiddenHost(this.host) {
		this.err = fmt.Errorf(`this hostname exsits on block list: %s`, this.host)
		return nil, this.err
	}

	this.addrs, this.err = this.resolver.LookupIPAddr(ctx, this.host)
	if len(this.addrs) == 0 {
		this.err = fmt.Errorf(`failed to lookup ip from host: %s`, this.host)
		return nil, this.err
	}

	for this.idx < len(this.addrs) {
		if this.ipRules.IsForbiddenIP(this.addrs[this.idx].IP) {
			this.err = fmt.Errorf(`this ip exists on block list: %s`, this.host)
			return nil, this.err
		}

		this.idx++
	}

	this.host = this.addrs[0].IP.String()
	addr = net.JoinHostPort(this.host, this.port)

	return this.parent.DialContext(ctx, network, addr)
}

func NewDialer(parent *net.Dialer, ip IPRules, host HostRules) ParanoidDialer {
	this := new(paranoidDialer)
	this.parent = parent
	this.ipRules = ip
	this.hostRules = host

	this.ip = nil
	this.host = ``
	this.port = ``
	this.addrs = nil
	this.err = nil
	this.idx = 0

	if parent.Resolver != nil {
		this.resolver = parent.Resolver
	} else {
		this.resolver = net.DefaultResolver
	}

	return this
}
