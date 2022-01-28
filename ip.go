package paranoid

import (
	"net"
	"sync"
)

type IPRules interface {
	IsForbiddenIP(addr net.IP) bool
}

type ipRules struct {
	mutex      *sync.Mutex
	matchers   []func(addr net.IP) bool
	idx        int
	length     int
	permission bool
}

func NewBlockIPRules(matchers ...func(addr net.IP) bool) IPRules {
	this := new(ipRules)
	this.permission = true
	this.matchers = matchers
	this.idx = 0
	this.length = len(matchers)
	this.mutex = new(sync.Mutex)
	return this
}

func NewAllowIPRules(matchers ...func(addr net.IP) bool) IPRules {
	this := new(ipRules)
	this.permission = false
	this.matchers = matchers
	this.idx = 0
	this.length = len(matchers)
	this.mutex = new(sync.Mutex)
	return this
}

func (this *ipRules) IsForbiddenIP(addr net.IP) bool {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	for this.idx < this.length {
		if this.matchers[this.idx](addr) {
			this.idx = 0
			return this.permission
		}

		this.idx++
	}

	this.idx = 0
	return !this.permission
}
