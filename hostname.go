package paranoid

import (
	"sync"
)

type Kind string

const (
	IsSameHost    Kind = "IsSameHost"
	HostHasPrefix Kind = "HostHasPrefix"
	HostHasSuffix Kind = "HostHasSuffix"
)

func (k Kind) String() string {
	return string(k)
}

type HostRule struct {
	kind  Kind
	addr  string
	idx   int
	mutex *sync.Mutex
}

func NewHostRule(kind Kind, addr string) *HostRule {
	this := new(HostRule)
	this.kind = kind
	this.addr = addr
	this.mutex = new(sync.Mutex)
	return this
}

func (this *HostRule) IsForbiddenHost(addr string) bool {
	switch this.kind {
	case IsSameHost:
		return this.addr == addr
	case HostHasPrefix:
		return this.hasPrefix(addr)
	case HostHasSuffix:
		return this.hasSuffix(addr)
	default:
		return false
	}
}

func (this *HostRule) Kind() Kind {
	return this.kind
}

func (this *HostRule) Addr() string {
	return this.addr
}

func (this *HostRule) hasPrefix(dest string) bool {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	if len(this.addr) > len(dest) {
		return false
	}

	for this.idx < len(this.addr) {
		if dest[this.idx] != this.addr[this.idx] {
			this.idx = 0
			return false
		}

		this.idx++
	}

	this.idx = 0
	return true
}

func (this *HostRule) hasSuffix(dest string) bool {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	if len(this.addr) > len(dest) {
		return false
	}

	this.idx = len(this.addr) - 1
	for this.idx >= 0 {
		if (dest[(len(dest)-len(this.addr))+this.idx]) != this.addr[this.idx] {
			return false
		}

		this.idx--
	}

	this.idx = 0
	return true
}
