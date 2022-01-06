package paranoid

import (
	"errors"
	"net"
	"sync"
)

var (
	ErrIPNetMissing = errors.New("list of net.IPNet is empty")
)

type IPBlocker struct {
	permission bool
	list       []*net.IPNet
	ip         *net.IPNet
	idx        int
	length     int
	result     bool
	mutex      *sync.Mutex
}

func NewIPBlocker(list []*net.IPNet, defaultPermittion bool) (*IPBlocker, error) {
	if len(list) == 0 {
		return nil, ErrIPNetMissing
	}

	this := new(IPBlocker)
	this.list = list
	this.permission = defaultPermittion
	this.idx = 0
	this.length = len(list)
	this.mutex = new(sync.Mutex)
	return this, nil
}

func (this *IPBlocker) IsForbiddenIP(ip net.IP) bool {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	for this.idx < this.length {
		if this.list[this.idx].Contains(ip) {
			this.idx = 0
			return this.permission
		}

		this.idx++
	}

	this.idx = 0
	return !this.permission
}
