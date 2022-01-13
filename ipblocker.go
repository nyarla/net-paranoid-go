package paranoid

import (
	"errors"
	"net"
	"sync"
)

var (
	// Errors
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

/*
  Instantiate IPBLocker.

    Arguments:

      - `list` is a list of *net.IPNet, it is used by detect IP address should be allowed or blocked.
      - `blockByDefalt` is a bit switch about IPBlocker should use []*net.IPNet as a allow list or block list.

    If set *true* to `blockByDefault`, IPBlocker blocks IP addresses by default
    and argument `list` means allow lists.

    Or if this value is *false*, IPBlocker allows IP addresses by default
    if IP address is not found inside ranges of []*net.IPNet.
    This mean argument `list` works as block list.
*/
func NewIPBlocker(list []*net.IPNet, blockByDefault bool) (*IPBlocker, error) {
	if len(list) == 0 {
		return nil, ErrIPNetMissing
	}

	this := new(IPBlocker)
	this.list = list
	this.permission = blockByDefault
	this.idx = 0
	this.length = len(list)
	this.mutex = new(sync.Mutex)
	return this, nil
}

/*
  This func check IP address should be blocked.

  If this func return true, you should be blocked this IP address.
  Or reutrn false, you should't.
*/
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
