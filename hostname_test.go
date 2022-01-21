package paranoid

import "testing"

func TestHostRule(t *testing.T) {
	tests := []struct {
		kind   Kind
		dest   string
		rule   string
		result bool
	}{
		{IsSameHost, `localhost`, `localhost`, true},
		{IsSameHost, `localhost`, `local`, false},

		{HostHasPrefix, `cdn.example.com`, `cdn`, true},
		{HostHasPrefix, `cdn.example.com`, `assets`, false},

		{HostHasSuffix, `cdn.example.com`, `example.com`, true},
		{HostHasSuffix, `assets.example.com`, `cdn.example.com`, false},
	}

	for _, test := range tests {
		if result := NewHostRule(test.kind, test.rule).IsForbiddenHost(test.dest); result != test.result {
			t.Errorf("failed to test:\nkind: %s\ndest: %s\nrule: %s\nwant: %v\nresult: %v", test.kind.String(), test.dest, test.rule, test.result, result)
		}
	}
}

func BenchmarkHostRule(b *testing.B) {
	var (
		hostname = `localhost`
		same     = NewHostRule(IsSameHost, `localhost`)
		prefix   = NewHostRule(HostHasPrefix, `cdn`)
		suffix   = NewHostRule(HostHasSuffix, `example.com`)
	)

	b.ReportAllocs()
	b.ResetTimer()

	for count := 0; count < b.N; count++ {
		same.IsForbiddenHost(hostname)
		prefix.IsForbiddenHost(hostname)
		suffix.IsForbiddenHost(hostname)
	}

}
