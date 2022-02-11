package main

func firstUniqChar(s string) byte {
	cmap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		value, ok := cmap[s[i]]
		if ok {
			cmap[s[i]] = value + 1
		} else {
			cmap[s[i]] = 1
		}
	}

	for i := 0; i < len(s); i++ {
		value, _ := cmap[s[i]]
		if value == 1 {
			return s[i]
		}
	}
	return ' '
}
