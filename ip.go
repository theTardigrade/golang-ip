package ip

import (
	"net"
	"net/http"
)

func Get(r *http.Request) (s string) {
	s = r.Header.Get("X-Real-Ip")

	if s == "" {
		s = r.Header.Get("X-Forwarded-For")
	}

	if s == "" {
		s = r.RemoteAddr
	}

	if s != "" {
		if p := net.ParseIP(s); p != nil {
			s = p.String()
		} else {
			s = ""
		}
	}

	return s
}
