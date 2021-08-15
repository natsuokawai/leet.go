package main

import "fmt"

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
	fmt.Println(lengthOfLastWord("luffy is still joyboy"))
}

func lengthOfLastWord(s string) int {
	l := 0
	prevIsSpace := false
	for _, c := range s {
		if c == ' ' {
			prevIsSpace = true
			continue
		} else {
			if prevIsSpace {
				l = 0
			}
			prevIsSpace = false
		}
		l++
	}
	return l
}
