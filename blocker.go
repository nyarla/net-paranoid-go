package paranoid

import (
	"net"
	"sync"
)

type IPRanges struct {
	permission bool
	ranges     []func(addr net.IP) bool
	idx        int
	length     int
	mutex      *sync.Mutex
}

func NewBlockIPRanges(ranges ...func(addr net.IP) bool) *IPRanges {
	this := new(IPRanges)
	this.permission = true
	this.ranges = ranges
	this.idx = 0
	this.length = len(ranges)
	this.mutex = new(sync.Mutex)
	return this
}

func NewAllowIPRanges(ranges ...func(addr net.IP) bool) *IPRanges {
	this := new(IPRanges)
	this.permission = false
	this.ranges = ranges
	this.idx = 0
	this.length = len(ranges)
	this.mutex = new(sync.Mutex)
	return this
}

func (this *IPRanges) IsForbiddenIP(addr net.IP) bool {
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
