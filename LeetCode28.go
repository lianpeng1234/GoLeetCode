package main

func strStr(haystack string, needle string) int {
	srcLen := len(haystack)
	targetLen := len(needle)

	if targetLen == 0 {
		return 0
	}

	if srcLen < targetLen {
		return -1
	}

	for i := 0; i < srcLen; i++ {
		if srcLen-targetLen < i {
			break
		}

		match := true
		for j := 0; j < targetLen; j++ {
			if haystack[i+j] == needle[j] {
				continue
			}
			match = false
			break
		}
		if match {
			return i
		}
	}

	return -1
}
