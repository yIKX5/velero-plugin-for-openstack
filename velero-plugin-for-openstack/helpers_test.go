package main

import (
	"testing"
)

func TestReplaceAccount(t *testing.T) {
	tests := []struct {
		name     string
		account  string
		endpoint string
		prefixes []string
		expected string
	}{
		{
			name:     "endpoint with account and existing prefix",
			account:  "john-doe",
			endpoint: "https://swift.openstack.domain.com:443/swift/v1/AUTH_randomstring1234",
			prefixes: []string{"AUTH_", "SERVICE_"},
			expected: "https://swift.openstack.domain.com:443/swift/v1/AUTH_john-doe",
		},
		{
			name:     "endpoint with account and non-existing prefix",
			account:  "john-doe",
			endpoint: "https://swift.openstack.domain.com:443/swift/v1/AUTH_randomstring1234",
			prefixes: []string{"IMAGE_"},
			expected: "https://swift.openstack.domain.com:443/swift/v1/AUTH_randomstring1234",
		},
		{
			name:     "endpoint without scheme, with account and existing prefix",
			account:  "john-doe",
			endpoint: "swift.domain.com:443/swift/v1/SERVICE_randomstring1234",
			prefixes: []string{"SERVICE_"},
			expected: "swift.domain.com:443/swift/v1/SERVICE_john-doe",
		},
	}
	for _, tt := range tests {
		path := ReplaceAccount(tt.account, tt.endpoint, tt.prefixes)
		if path != tt.expected {
			t.Errorf("[%s] failed - output %s doesn't match expected '%s'", tt.name, path, tt.expected)
		}
	}
}
