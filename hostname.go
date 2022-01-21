package paranoid

import (
	"regexp"
	"strings"
	"sync"
)

func IsSameHost(src, cmp string) bool {
	return src == cmp
}

func HostPrefix(src, cmp string) bool {
	return strings.HasPrefix(src, cmp)
}

func HostSuffix(src, cmp string) bool {
	return strings.HasSuffix(src, cmp)
}

type HostMatcher interface {
	MatchHost(src string) bool
}

type stringHostMatcher struct {
	cmp     string
	matcher func(src, cmp string) bool
}

func (this *stringHostMatcher) MatchHost(src string) bool {
	return this.matcher(src, this.cmp)
}

func StringHostMatcher(cmp string, matcher func(src, cmp string) bool) HostMatcher {
	this := new(stringHostMatcher)
	this.cmp = cmp
	this.matcher = matcher
	return this
}

type regexHostMatcher struct {
	re *regexp.Regexp
}

func (this *regexHostMatcher) MatchHost(src string) bool {
	return this.re.MatchString(src)
}

func RegexpHostMatcher(re *regexp.Regexp) HostMatcher {
	this := new(regexHostMatcher)
	this.re = re
	return this
}

type HostRules interface {
	IsForbiddenHost(addr string) bool
}

type hostRules struct {
	permission bool
	matchers   []HostMatcher
	idx        int
	length     int
	mutex      *sync.Mutex
}

func NewBlockHostRules(matchers ...HostMatcher) HostRules {
	this := new(hostRules)
	this.permission = true
	this.matchers = matchers
	this.idx = 0
	this.length = len(matchers)
	this.mutex = new(sync.Mutex)

	return this
}

func NewAllowHostRules(matchers ...HostMatcher) HostRules {
	this := new(hostRules)
	this.permission = false
	this.matchers = matchers
	this.idx = 0
	this.length = len(matchers)
	this.mutex = new(sync.Mutex)

	return this
}

func (this *hostRules) IsForbiddenHost(src string) bool {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	for this.idx < this.length {
		if this.matchers[this.idx].MatchHost(src) {
			this.idx = 0
			return this.permission
		}

		this.idx++
	}

	this.idx = 0
	return !this.permission
}
