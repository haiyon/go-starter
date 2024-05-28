package util

import (
	"fmt"
	"go-starter/internal/config"
)

// Find returns the smallest index i at which x == a[i],
// or len(a) if there is no such index.
func Find(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return len(a)
}

func FindID(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return len(a)
}

// Contains tells whether a contains x.
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

// GetDomain Get the run domain
func GetDomain(domain string) string {
	if domain == "localhost" {
		return fmt.Sprintf("%v://%v:%d", config.G.Protocol, config.G.Domain, config.G.Port)
	}
	return fmt.Sprintf("%v://%v", config.G.Protocol, config.G.Domain)
}