package cmd

import (
	"testing"
)

func TestParseURL(t *testing.T) {
	tests := []struct {
		url      string
		expected [3]string
	}{
		{"https://github.com/organisation/repository.git", [3]string{"github.com", "organisation", "repository"}},
		{"git@github.com:organisation/repository.git", [3]string{"github.com", "organisation", "repository"}},
	}

	for _, test := range tests {
		hoster, org, repo := parseURL(test.url)
		if hoster != test.expected[0] || org != test.expected[1] || repo != test.expected[2] {
			t.Errorf("parseURL(%s) = %s, %s, %s; want %s, %s, %s", test.url, hoster, org, repo, test.expected[0], test.expected[1], test.expected[2])
		}
	}
}
