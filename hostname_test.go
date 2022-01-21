package paranoid

import (
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"testing"
)

func TestHostMatchers(t *testing.T) {
	tests := []struct {
		matcher func(src, cmp string) bool
		src     string
		cmp     string
		result  bool
	}{
		{IsSameHost, `localhost`, `localhost`, true},
		{IsSameHost, `localhost`, `example.com`, false},

		{HostPrefix, `cdn.example.com`, `cdn`, true},
		{HostPrefix, `assets.example.com`, `cdn`, false},

		{HostSuffix, `cdn.example.com`, `example.com`, true},
		{HostSuffix, `cdn.example.com`, `net`, false},
	}

	for _, test := range tests {
		fname := strings.Split(
			(runtime.FuncForPC(reflect.ValueOf(test.matcher).Pointer()).Name()), `.`,
		)

		matcher := StringHostMatcher(test.cmp, test.matcher)

		if result := matcher.MatchHost(test.src); result != test.result {
			t.Errorf("Test %s failed:\nsrc: %s\ncmp: %s\nwant: %v\nresult: %v",
				fname[len(fname)-1],
				test.src, test.cmp, test.result, result,
			)
		}
	}

	re := RegexpHostMatcher(regexp.MustCompile(`^localhost$`))
	if result := re.MatchHost(`localhost`); result != true {
		t.Errorf("Test RegexpHostMatcher.MatchHost failed:\nsrc: localhost\ncmp: ^localhost$\nwant: %v\nresult: %v",
			true, result,
		)
	}
}

func TestHostRules(t *testing.T) {
	blocker := NewBlockHostRules(
		StringHostMatcher(`localhost`, IsSameHost),
		StringHostMatcher(`localhost`, HostPrefix),
		StringHostMatcher(`localhost`, HostSuffix),
	)

	allower := NewAllowHostRules(
		StringHostMatcher(`localhost`, IsSameHost),
		StringHostMatcher(`localhost`, HostPrefix),
		StringHostMatcher(`localhost`, HostSuffix),
	)

	if !blocker.IsForbiddenHost(`localhost`) {
		t.Errorf("Test HostRules.IsForbiddenHost failed:\nwant: true\ngot: false")
	}

	if allower.IsForbiddenHost(`localhost`) {
		t.Errorf("Test HostRules.IsForbiddenHost failed:\nwant: false\ngot: true")
	}
}

func BenchmarkHostRules(b *testing.B) {
	allower := NewAllowHostRules(
		StringHostMatcher(`localhost`, IsSameHost),
		StringHostMatcher(`localhost`, HostPrefix),
		StringHostMatcher(`localhost`, HostSuffix),
		RegexpHostMatcher(regexp.MustCompile(`^localhost$`)),
	)

	hostname := `localhost`

	b.ReportAllocs()
	b.ResetTimer()

	for idx := 0; idx < b.N; idx++ {
		allower.IsForbiddenHost(hostname)
	}
}
