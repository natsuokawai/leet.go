package main

import "fmt"

func main() {
	fmt.Println(strStr("hello", "ll"))
	fmt.Println(strStr("aaa", "bb"))
	fmt.Println(strStr("aa", "aabb"))
	fmt.Println(strStr("", ""))
}

func strStr(haystack string, needle string) int {
	hlen := len(haystack)
	nlen := len(needle)

	if nlen == 0 {
		return 0
	}
	if hlen == 0 {
		return -1
	}
	if hlen < nlen {
		return -1
	}

	for i := 0; i <= hlen-nlen; i++ {
		if haystack[i:i+nlen] == needle {
			return i
		}
	}
	return -1
}
