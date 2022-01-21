package paranoid

import (
	"net"
	"sync"
)

type IPRules struct {
	permission bool
	ranges     []func(addr net.IP) bool
	idx        int
	length     int
	mutex      *sync.Mutex
}

func NewBlockIPRules(ranges ...func(addr net.IP) bool) *IPRules {
	this := new(IPRules)
	this.permission = true
	this.ranges = ranges
	this.idx = 0
	this.length = len(ranges)
	this.mutex = new(sync.Mutex)
	return this
}

func NewAllowIPRules(ranges ...func(addr net.IP) bool) *IPRules {
	this := new(IPRules)
	this.permission = false
	this.ranges = ranges
	this.idx = 0
	this.length = len(ranges)
	this.mutex = new(sync.Mutex)
	return this
}

func (this *IPRules) IsForbiddenIP(addr net.IP) bool {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	for this.idx < this.length {
		if this.ranges[this.idx](addr) {
			this.idx = 0
			return this.permission
		}

		this.idx++
	}

	this.idx = 0
	return !this.permission
}
