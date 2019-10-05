package ip

import (
	"net/http"
	"strings"
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
		splitS := strings.Split(s, "::")

		if splitSLastIndex := len(splitS) - 1; splitSLastIndex > 0 {
			splitSplitS := strings.Split(splitS[splitSLastIndex], ":")

			if splitSplitSLastIndex := len(splitSplitS) - 1; splitSplitSLastIndex > 0 {
				splitSplitS = splitSplitS[:splitSplitSLastIndex]
			}

			splitS[splitSLastIndex] = strings.Join(splitSplitS, ":")
		}

		s = strings.Join(splitS, "::")
	}

	return s
}
