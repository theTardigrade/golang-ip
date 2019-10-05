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
		splitSLen := len(splitS)
		finalSplitSIndex := splitSLen - 1
		finalSplitS := splitS[finalSplitSIndex]
		splitFinalSplitS := strings.Split(finalSplitS, ":")
		if l := len(splitFinalSplitS); l > 1 {
			splitFinalSplitS = splitFinalSplitS[:l-1]
		}
		splitS = append(splitS[:finalSplitSIndex], strings.Join(splitFinalSplitS, ":"))
		s = strings.Join(splitS, "::")
	}

	return s
}
