package paranoid

import (
	"errors"
	"net"
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
	return this, nil
}

func (this *IPBlocker) IsForbiddenIP(ip net.IP) bool {
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
