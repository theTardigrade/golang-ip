package ip

import (
	"net/http"
	"strings"
)

func removePortIPv4(s string) string {
	splitS := strings.Split(s, ":")

	if splitSLastIndex := len(splitS) - 1; splitSLastIndex > 0 {
		splitS = splitS[:splitSLastIndex]
	}

	return strings.Join(splitS, ":")
}

func removePortIPv6(s string) string {
	splitS := strings.Split(s, "::")

	if splitSLastIndex := len(splitS) - 1; splitSLastIndex >= 0 {
		splitS[splitSLastIndex] = removePortIPv4(splitS[splitSLastIndex])
	}

	return strings.Join(splitS, "::")
}

// Get returns a string representation of the client's IP address
func Get(r *http.Request) (s string) {
	s = r.Header.Get("X-Real-Ip")

	if s == "" {
		s = r.Header.Get("X-Forwarded-For")
	}

	if s == "" {
		s = r.RemoteAddr
	}

	if s != "" { // remove port
		if strings.Contains(s, "::") {
			s = removePortIPv6(s)
		} else {
			s = removePortIPv4(s)
		}
	}

	return s
}
