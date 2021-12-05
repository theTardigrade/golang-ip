package ip

import "strings"

func portRemoveIPv4(s string) string {
	splitS := strings.Split(s, ":")

	if splitSLastIndex := len(splitS) - 1; splitSLastIndex > 0 {
		splitS = splitS[:splitSLastIndex]
	}

	return strings.Join(splitS, ":")
}

func portRemoveIPv6(s string) string {
	splitS := strings.Split(s, "::")

	if splitSLastIndex := len(splitS) - 1; splitSLastIndex >= 0 {
		splitS[splitSLastIndex] = portRemoveIPv4(splitS[splitSLastIndex])
	}

	return strings.Join(splitS, "::")
}
