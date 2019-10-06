package ip

import "strings"

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
